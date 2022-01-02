package controllers

import (
	"encoding/json"
	"xcloud/util"
	"context"
	//"strconv"
	"xcloud/models"
	//"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/orm"
)

type ConfigmapController struct {
	beego.Controller
}


// @router / [post]
func (this *ConfigmapController) Post() {
	var cm map[string]interface{}
	json.Unmarshal([]byte(this.Ctx.Input.RequestBody), &cm)
	beego.Info(cm)
	cluster_id :=  cm["cluster"]
	namespace := cm["namespace"].(string)
	o := orm.NewOrm()
	var cluster models.Cluster
	o.QueryTable("Cluster").Filter("Id",cluster_id).One(&cluster)
	clustername := cluster.Name
	clientset := util.Getclient(clustername)
	dataMap := make(map[string]string)
	for _,v := range cm["keyvalue"].([]interface{}) {
		dataMap[v.(map[string]interface{})["key"].(string)] = v.(map[string]interface{})["value"].(string)
	}
	beego.Info(dataMap)
	configmap_yaml := &apiv1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind: "configmap",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: cm["configmapname"].(string),
		},
		Data: dataMap,
	}
	_, err3 := clientset.CoreV1().ConfigMaps(namespace).Create(context.TODO(), configmap_yaml, metav1.CreateOptions{})
	beego.Info(err3)
	if err3 != nil {
		this.CustomAbort(500 ,err3.Error())
	}
}

// @router / [get]
func (this *ConfigmapController) Get(){
	o := orm.NewOrm()
	var clusters []orm.ParamsList
	o.QueryTable("Cluster").ValuesList(&clusters, "Id","Name")
	beego.Info(clusters,"ppppppppp")
}


// @router / [get]
func (this *ConfigmapController) GetNamespace() {
	var cluster_id = this.GetString("cluster_id")
	json.Unmarshal([]byte(this.Ctx.Input.RequestBody), &cluster_id)
	beego.Info(cluster_id)
	o := orm.NewOrm()
	var cluster models.Cluster
	o.QueryTable("Cluster").Filter("Id", cluster_id).One(&cluster)
	clustername := cluster.Name
	clientset := util.Getclient(clustername)
	NameSpaveList, _ := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	this.Data["json"] = NameSpaveList
	this.ServeJSON()
}