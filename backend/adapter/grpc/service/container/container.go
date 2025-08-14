package container

import (
	ifaceController "ensina-renda/controller/iface"
)

type EnsinaRendaContainerInterface interface {
	UsuarioController() ifaceController.UsuarioController
	AulaController() ifaceController.AulaController
	ModuloController() ifaceController.ModuloController
}

type EnsinaRendaContainer struct {
	usuarioController ifaceController.UsuarioController
	aulaController    ifaceController.AulaController
	moduloController  ifaceController.ModuloController
}

func NewEnsinaRendaContainer(
	usuarioController ifaceController.UsuarioController,
	aulaController ifaceController.AulaController,
	moduloController ifaceController.ModuloController,
) *EnsinaRendaContainer {
	return &EnsinaRendaContainer{
		usuarioController: usuarioController,
		aulaController:    aulaController,
		moduloController:  moduloController,
	}
}

func (c *EnsinaRendaContainer) UsuarioController() ifaceController.UsuarioController {
	return c.usuarioController
}

func (c *EnsinaRendaContainer) AulaController() ifaceController.AulaController {
	return c.aulaController
}

func (c *EnsinaRendaContainer) ModuloController() ifaceController.ModuloController {
	return c.moduloController
}
