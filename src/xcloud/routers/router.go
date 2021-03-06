// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"xcloud/controllers"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego"
)

func init() {

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "x-token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/cluster",
			beego.NSInclude(
				&controllers.ClusterController{},
			),
		),
		beego.NSNamespace("/clusterdetail",
			beego.NSInclude(
				&controllers.ClusterdetailController{},
			),
		),
		beego.NSNamespace("/clusterresdetail",
			beego.NSInclude(
				&controllers.ClusterresdetailController{},
			),
		),
		beego.NSNamespace("/env",
		    beego.NSInclude(
			    &controllers.EnvController{},
		    ),
	    ),
		beego.NSNamespace("/registry",
			beego.NSInclude(
				&controllers.RegistryController{},
			),
		),
		beego.NSNamespace("/cert",
			beego.NSInclude(
				&controllers.CertController{},
			),
		),
		beego.NSNamespace("/clusterimages",
			beego.NSInclude(
				&controllers.ImageController{},
			),
		),
		beego.NSNamespace("/imagetaglist",
			beego.NSRouter(
				"/",&controllers.ImageController{},"get:GetImageTag",
			),
		),
		beego.NSNamespace("/configmap",
			beego.NSInclude(
				&controllers.ConfigmapController{},
			),
		),
		beego.NSNamespace("/namespace",
			beego.NSRouter(
				"/",&controllers.ConfigmapController{},"get:GetNamespace",
			),
		),
		beego.NSNamespace("/containerwebshell",
			beego.NSRouter(
				"/",&controllers.ContainerWebshellController{},"get:WsHandler",
			),
		),
		beego.NSNamespace("/buildlog",
			beego.NSRouter(
				"/",&controllers.ProjectbuildController{},"get:Ws",
			),
		),
		beego.NSNamespace("/getbuildresult",
			beego.NSRouter(
				"/",&controllers.ProjectbuildController{},"get:GetBuildResult",
			),
		),
	)
	//web.AddNamespace(ns)
	beego.AddNamespace(ns)
}
