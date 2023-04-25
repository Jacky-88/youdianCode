package service

import (
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/release"
	"io"
	"k8s-demo/config"
	"k8s-demo/dao"
	"k8s-demo/model"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

var HelmStore helmStore

type helmStore struct{}

type releaseElement struct {
	Name         string `json:"name"`
	Namespace    string `json:"namespace"`
	Revision     string `json:"revision"`
	Updated      string `json:"updated"`
	Status       string `json:"status"`
	Chart        string `json:"chart"`
	ChartVersion string `json:"chart_version"`
	AppVersion   string `json:"app_version"`

	Notes string `json:"notes,omitempty"`
}
type releaseElements struct {
	Items []*releaseElement `json:"items"`
	Total int                `json:"total"`
}
//release列表
func(*helmStore) ListReleases(actionConfig *action.Configuration, filterName string) (*releaseElements, error) {
	//startSet := (page-1) * limit
	client := action.NewList(actionConfig)
	client.Filter = filterName
	client.All = true
	//client.Limit = limit
	//client.Offset = startSet
	client.TimeFormat = "2006-01-02 15:04:05"
	//是否已经部署
	client.Deployed = true
	results, err := client.Run()
	if err != nil {
		logger.Error("listReleases请求失败, " + err.Error())
		return nil, errors.New("listReleases请求失败, " + err.Error())
	}
	total := len(results)
	elements := make([]*releaseElement, 0, len(results))
	for _, r := range results {
		elements = append(elements, constructReleaseElement(r, false))
	}
	return &releaseElements{
		Items: elements,
		Total: total,
	}, nil
}

//release详情
func(*helmStore) DetailRelease(actionConfig *action.Configuration, release string) (*release.Release, error) {
	client := action.NewGet(actionConfig)
	data, err := client.Run(release)
	if err != nil {
		logger.Error("detailRelease请求失败, " + err.Error())
		return nil, errors.New("detailRelease请求失败, " + err.Error())
	}
	return data, nil
}

//release安装
func(*helmStore) InstallRelease(actionConfig *action.Configuration, release, chart, namespace string) error {
	client := action.NewInstall(actionConfig)
	client.ReleaseName = release
	client.Namespace = namespace

	splitChart := strings.Split(chart, ".")
	if splitChart[len(splitChart)-1] == "tgz" && !strings.Contains(chart, ":") {
		chart = config.UploadPath + "/" + chart
	}
	chartRequested, err := loader.Load(chart)
	if err != nil {
		logger.Error("加载chart文件失败, " + err.Error())
		return errors.New("加载chart文件失败, " + err.Error())
	}
	vals := map[string]interface{}{}
	_, err = client.Run(chartRequested, vals)
	if err != nil {
		logger.Error("安装release失败, " + err.Error())
		return errors.New("安装release失败, " + err.Error())
	}
	return nil
}

//release卸载
func(*helmStore) UninstallRelease(actionConfig *action.Configuration, release, namespace string) error {
	client := action.NewUninstall(actionConfig)
	_, err := client.Run(release)
	if err != nil {
		logger.Error("卸载release失败, " + err.Error())
		return errors.New("卸载release失败, " + err.Error())
	}
	return nil
}

//chart文件上传
func(*helmStore) UploadChartFile(file multipart.File, header *multipart.FileHeader) error {
	filename := header.Filename
	t := strings.Split(filename, ".")
	if t[len(t)-1] != "tgz" {
		logger.Error("chart文件必须以.tgz结尾")
		return errors.New("chart文件必须以.tgz结尾")
	}
	filePath := config.UploadPath + "/" + filename
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		logger.Error("chart文件已存在")
		return errors.New("chart文件已存在")
	}
	out, err := os.Create(filePath)
	if err != nil {
		logger.Error("创建chart文件失败, " + err.Error())
		return errors.New("创建chart文件失败, " + err.Error())
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		logger.Error("创建chart文件失败2, " + err.Error())
		return errors.New("创建chart文件失败2, " + err.Error())
	}
	return nil
}

//chart文件删除
func(*helmStore) DeleteChartFile(chart string) error {
	filePath := config.UploadPath + "/" + chart
	// not exist,ok
	_, err := os.Stat(filePath)
	if err != nil || os.IsNotExist(err) {
		logger.Error("chart文件不存在, " + err.Error())
		return errors.New("chart文件不存在, " + err.Error())
	}
	err = os.Remove(filePath)
	if err != nil {
		logger.Error("chart文件删除失败, " + err.Error())
		return errors.New("chart文件删除失败, " + err.Error())
	}
	return nil
}

//chart列表
func(*helmStore) ListCharts(name string, page, limit int) (*dao.Charts, error) {
	return dao.Chart.GetList(name, page, limit)
}

//chart新增
func(*helmStore) AddChart(chart *model.Chart) error {
	_, has, err := dao.Chart.Has(chart.Name)
	if err != nil {
		return err
	}
	if has {
		return errors.New("该数据已存在，请重新添加")
	}
	if err := dao.Chart.Add(chart); err != nil {
		return err
	}
	return nil
}

//Chart更新
func(h *helmStore) UpdateChart(chart *model.Chart) error {
	oldChart, _, err := dao.Chart.Has(chart.Name)
	if err != nil {
		return err
	}
	fmt.Println(chart.FileName, oldChart.FileName)
	if chart.FileName != "" && chart.FileName != oldChart.FileName {
		err = h.DeleteChartFile(oldChart.FileName)
		if err != nil {
			return err
		}
	}
	return dao.Chart.Update(chart)
}

//Chart删除
func(h *helmStore) DeleteChart(chart *model.Chart) error {
	//删除文件
	err := h.DeleteChartFile(chart.FileName)
	if err != nil {
		return err
	}
	//删除数据
	return dao.Chart.Delete(chart.ID)
}

//release内容过滤
func constructReleaseElement(r *release.Release, showStatus bool) *releaseElement {
	element := &releaseElement{
		Name:         r.Name,
		Namespace:    r.Namespace,
		Revision:     strconv.Itoa(r.Version),
		Status:       r.Info.Status.String(),
		Chart:        r.Chart.Metadata.Name,
		ChartVersion: r.Chart.Metadata.Version,
		AppVersion:   r.Chart.Metadata.AppVersion,
	}
	if showStatus {
		element.Notes = r.Info.Notes
	}
	t := "-"
	if tspb := r.Info.LastDeployed; !tspb.IsZero() {
		t = tspb.String()
	}
	element.Updated = t

	return element
}