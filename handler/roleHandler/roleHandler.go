package roleHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"usercenter/models"
	"usercenter/models/sysRole"
	"usercenter/validator"
)

// 角色列表
// @Summary 角色列表
// @Description 角色列表
// @Tags 角色
// @Success 200 {object} models.PagedResponse{data=[]sysRole.SysRole}
// @param pageIndex query number false "pageIndex" default(1)
// @param pageSize query number false "pageSize" default(10)
// @param keyWord query string false "关键词"
// @Router /roles [get]
func RoleList(c *gin.Context) {
	var req = sysRole.GetAllRoleRequest{}
	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	req.PageIndex = pageIndex
	req.PageSize = pageSize
	req.KeyWord = c.DefaultQuery("keyWord", "")

	list, count := sysRole.GetAllRole(req)
	var res models.PagedResponse
	res.Total = count
	res.Data = list

	c.JSON(http.StatusOK, res)
}

// 添加角色
// @Summary 添加角色
// @Description 角色列表
// @Tags 角色
// @Success 200 {string} string ""
// @param role body sysRole.AddRoleRequest true "添加角色"
// @Router /roles [post]
func AddRole(c *gin.Context) {
	var req = sysRole.AddRoleRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: validator.Translate(err)})
		return
	}

	if err := sysRole.AddRole(req); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.String(http.StatusOK, "")
}

// 根据id获取角色信息
// @Summary 根据id获取角色信息
// @Description 根据id获取角色信息
// @Tags 角色
// @Success 200 {object} sysRole.SysRole
// @param id path number false "id"
// @Router /roles/{id} [get]
func GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	role := sysRole.GetById(id)
	c.JSON(http.StatusOK, role)
}

// 根据id更新角色
// @Summary 根据id更新角色
// @Description 根据id更新角色
// @Tags 角色
// @Success 200 {string} string ""
// @param id path number false "id"
// @param body body sysRole.UpdateRoleRequest true "更新角色"
// @Router /roles/{id} [put]
func UpdateRole(c *gin.Context) {
	req := sysRole.UpdateRoleRequest{}
	req.Id, _ = strconv.ParseInt(c.Param("id"), 10, 64)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: validator.Translate(err)})
		return
	}

	if err := sysRole.UpdateRole(req); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.String(http.StatusOK, "")
}

// 根据id删除角色
// @Summary 根据id删除角色
// @Description 根据id更新角色
// @Tags 角色
// @Success 200 {string} string ""
// @param id path number false "id"
// @Router /roles/{id} [delete]
func DeleteById(c *gin.Context) {
	role := sysRole.SysRole{}
	role.Id, _ = strconv.ParseInt(c.Param("id"), 10, 64)

	if err := role.DeleteById(); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.String(http.StatusOK, "")
}
