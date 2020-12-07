package userHandler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"usercenter/models"
	"usercenter/models/sysRole"
	"usercenter/models/sysUser"
	"usercenter/validator"
)

// 用户列表
// @Summary 用户列表
// @Description 用户列表
// @Tags 用户
// @Success 200 {object} models.PagedResponse{data=[]sysUser.SysUser}
// @param pageIndex query number false "pageIndex" default(1)
// @param pageSize query number false "pageSize" default(10)
// @param keyWord query string false "关键词"
// @Router /users [get]
func UserList(c *gin.Context) {
	var req = sysUser.GetAllUserRequest{}
	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	req.PageIndex = pageIndex
	req.PageSize = pageSize
	req.KeyWord = c.DefaultQuery("keyWord", "")
	req.Username = c.DefaultQuery("username", "")

	users, count := sysUser.GetAllUser(req)

	var res models.PagedResponse
	res.Total = count
	res.Data = users

	c.JSON(http.StatusOK, res)
}

// 根据id获取用户信息
// @Summary 根据id获取用户信息
// @Description 根据id获取用户信息
// @Tags 用户
// @Success 200 {object} sysUser.SysUser
// @param id path number false "id"
// @Router /users/{id} [get]
func GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panicln("get user by id no found id", err.Error())
	}
	user := sysUser.GetById(id)
	c.JSON(http.StatusOK, user)
}

// 根据id更新用户
// @Summary 根据id更新用户
// @Description 根据id更新用户
// @Tags 用户
// @Success 200 {string} string ""
// @param id path number false "id"
// @param body body sysUser.UpdateUserRequest true "更新用户"
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	req := sysUser.UpdateUserRequest{}
	req.Id, _ = strconv.Atoi(c.Param("id"))

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	success := sysUser.UpdateUser(req)
	if !success {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "更新失败"})
		return
	}
	c.String(http.StatusOK, "")
}

// 添加用户
// @Summary 添加用户
// @Description 角色用户
// @Tags 用户
// @Success 200 {string} string ""
// @param role body sysUser.AddUserRequest true "添加用户"
// @Router /users [post]
func AddUser(c *gin.Context) {
	req := sysUser.AddUserRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: validator.Translate(err)})
		return
	}
	req.RegIp = c.ClientIP()
	success, err := sysUser.AddUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "添加用户失败"})
		return
	}
	c.String(http.StatusOK, "")
}

// 删除用户
// @Summary 删除用户
// @Description 删除用户
// @Tags 用户
// @Success 200 {string} string ""
// @param id path number false "id"
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	user := sysUser.SysUser{}
	user.Id, _ = strconv.Atoi(c.Param("id"))

	if err := user.LogicalDeleteById(); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.String(http.StatusOK, "")
}

// 更新用户密码
// @Summary 更新用户密码
// @Description 更新用户密码
// @Tags 用户
// @Success 200 {string} string ""
// @param id path number false "id"
// @param body body sysUser.UpdatePasswordByIdRequest true "更新用户密码"
// @Router /users/{id}/attr/password [put]
func UpdatePassword(c *gin.Context) {
	request := sysUser.UpdatePasswordByIdRequest{}
	request.Id, _ = strconv.Atoi(c.Param("id"))
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	if err := sysUser.UpdatePasswordById(request); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.String(http.StatusOK, "")
}

// 获取用户角色
// @Summary 获取用户角色
// @Description 获取用户角色
// @Tags 用户
// @Success 200 {object} sysRole.RoleWithChecked
// @param id path number false "id"
// @Router /users/{id}/roles [put]
func GetAllRoleWithCheckedByUserId(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	list := sysRole.GetAllRoleWithCheckedByUserId(userId)

	c.JSON(http.StatusOK, list)
}

// 更新用户角色
// @Summary 更新用户角色
// @Description 更新用户角色
// @Tags 用户
// @Success 200 {string} string ""
// @param id path number false "id"
// @param body body sysUser.UpdatePasswordByIdRequest true "更新用户密码"
// @Router /users/{id}/roles [put]
func UpdateUserRole(c *gin.Context) {
	var json sysUser.UpdateUserRoleRequest
	userId, _ := strconv.Atoi(c.Param("id"))
	json.Id = userId
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	sysUser.UpdateUserRole(json)
	c.String(http.StatusOK, "")
}
