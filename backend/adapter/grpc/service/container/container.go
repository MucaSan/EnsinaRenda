package container

import (
	ifaceController "ensina-renda/controller/iface"
)

type EnsinaRendaContainerInterface interface {
	UsuarioController() ifaceController.UsuarioController
	AulaController() ifaceController.AulaController
}

type EnsinaRendaContainer struct {
	usuarioController ifaceController.UsuarioController
	aulaController    ifaceController.AulaController
}

func NewEnsinaRendaContainer(
	usuarioController ifaceController.UsuarioController,
	aulaController ifaceController.AulaController,
) *EnsinaRendaContainer {
	return &EnsinaRendaContainer{
		usuarioController: usuarioController,
		aulaController:    aulaController,
	}
}

func (c *EnsinaRendaContainer) UsuarioController() ifaceController.UsuarioController {
	return c.usuarioController
}

func (c *EnsinaRendaContainer) AulaController() ifaceController.AulaController {
	return c.aulaController
}
