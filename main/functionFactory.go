package main

// // GetFunctionFactory ...
// func GetFunctionFactory() *model.FunctionFactory {
// 	return &model.FunctionFactory{
// 		NewDataAccessLayer: dependency.NewDataAccessLayer,
// 		NewAnotherDal:      dependency.NewAnotherDal,
// 	}
// }

// // GetFactoryNewDataAccessLayer ...
// func GetFactoryNewDataAccessLayer() model.FactoryNewDataAccessLayer {
// 	return model.FactoryNewDataAccessLayer{
// 		NewDataAccessLayer: dependency.NewDataAccessLayer,
// 	}
// }

// // GetFactoryNewAnotherDal ...
// func GetFactoryNewAnotherDal() model.FactoryNewAnotherDal {
// 	return model.FactoryNewAnotherDal{
// 		NewAnotherDal: dependency.NewAnotherDal,
// 	}
// }

// func GetAllFactory() model.AllFactory {
// 	return model.AllFactory{
// 		FactoryNewDataAccessLayer: GetFactoryNewDataAccessLayer(),
// 		FactoryNewAnotherDal:      GetFactoryNewAnotherDal(),
// 	}
// }
