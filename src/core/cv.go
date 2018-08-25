package core

func CrossValidate(algorithm Algorithm, dataSet TrainSet, metrics []Metrics, cv int, seed int64,
	options ...OptionSetter) [][]float64 {
	ret := make([][]float64, len(metrics))
	for i := 0; i < len(ret); i++ {
		ret[i] = make([]float64, cv)
	}
	// Split data set
	trainFolds, testFolds := dataSet.KFold(cv, seed)
	for i := 0; i < cv; i++ {
		trainFold := trainFolds[i]
		testFold := testFolds[i]
		algorithm.Fit(trainFold, options...)
		predictions := make([]float64, testFold.Length())
		for j := 0; j < testFold.Length(); j++ {
			userId := testFold.Users()[j]
			itemId := testFold.Items()[j]
			predictions[j] = algorithm.Predict(userId, itemId)
		}
		truth := testFold.Ratings()
		// Metrics
		for j := 0; j < len(ret); j++ {
			ret[j][i] = metrics[j](predictions, truth)
		}
	}
	return ret
}
