package k8s

import (
	"context"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"question.com/global"
	"question.com/model/common/response"
	ns_res "question.com/model/ns/response"
)

type NsApi struct {
}

func (*NsApi) GetNsList(c *gin.Context) {

	ctx := context.TODO()
	list, err := global.KubeConfigset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	nsList := make([]ns_res.Ns, 0)

	for _, item := range list.Items {
		nsList = append(nsList, ns_res.Ns{
			Name:              item.Name,
			CreationTimestamp: item.CreationTimestamp.Unix(),
			Status:            string(item.Status.Phase),
		})
	}
	response.SuccessWithDetailed(c, nsList, "获取成功")
}
