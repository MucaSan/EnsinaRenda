package container

import (
	ifaceController "ensina-renda/controller/iface"
)

type EnsinaRendaContainerInterface interface {
	UsuarioController() ifaceController.UsuarioController
}

type EnsinaRendaContainer struct {
	usuarioController ifaceController.UsuarioController
}

func NewEnsinaRendaContainer(usuarioController ifaceController.UsuarioController) *EnsinaRendaContainer {
	return &EnsinaRendaContainer{
		usuarioController: usuarioController,
	}
}

func (c *EnsinaRendaContainer) UsuarioController() ifaceController.UsuarioController {
	return c.usuarioController
}
