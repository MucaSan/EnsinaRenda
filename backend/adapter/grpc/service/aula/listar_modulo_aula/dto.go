package listar_modulo_aula

import (
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/domain/model"
	"strconv"
)

func ConverterMapaParaPb(mapaModuloAula map[int][]*model.UsuarioModuloAula) []*pb.ModuloAulaAluno {
	mapModuloPb := map[int][]*pb.Aula{}
	moduloAulas := []*pb.ModuloAulaAluno{}

	for idModulo, aulas := range mapaModuloAula {
		for _, aula := range aulas {
			mapModuloPb[idModulo] = append(mapModuloPb[idModulo], &pb.Aula{
				Id:     strconv.Itoa(aula.IDAula),
				Status: aula.Concluido,
			})
		}
	}

	for pbIdModulo, pbAulas := range mapModuloPb {
		moduloAulas = append(
			moduloAulas, &pb.ModuloAulaAluno{
				IdModulo: strconv.Itoa(pbIdModulo),
				Aulas:    pbAulas,
			})
	}

	return moduloAulas
}
