from pandas import datetime
from matplotlib import pyplot
from statsmodels.tsa.arima_model import ARIMA
from pandas import DataFrame
import matplotlib.pyplot as plt
import numpy as np
from random import randint
times = []
values = []
sum = 0
index = 0


time_list = []
value_list = []
with open("ElectricityByDay.csv") as input_:
    for line in input_:
        time_list.append(line.split(",")[0][-2])
        value_list.append(float(line.split(",")[1][1:-2].strip("\n")))


def get_svm_index_based_method1(index, time_list, value_list):
    """ used to get index in method1 mentioned in doc of svm_predict()
    """
    data = np.array(time_list[index:index+7]).T.reshape(-1, 1)
    label = np.array(value_list[index:index+7])
    return (data, label)


def get_svm_index_based_method2(index, time_list, value_list):
    """ used to get index in method2 mentioned in doc of svm_predict()
    """
    data = np.array(list(time_list[index+7*i] for i in range(7))).T.reshape(-1, 1)
    label = np.array(list([value_list[index+7*i] for i in range(7)])).T.reshape(-1, 1)
    return (data,label)


def svm_predict():

    # method 1 : past 7 days                         called month on month
    # method 2 : the same day in the past 7 weeks   called year on year
    # method 3 : the same day in the past 7 years   # very limited data , so we will not adopt this method
    # For method 1
    # firstly , get 200 records
    # then get random value
    expected_result_list1 = []
    predict_result_list1 = []
    expected_result_list2 = []
    predict_result_list2 = []
    from sklearn.svm import SVR
    for _ in range(200):
        index1 = randint(1, 1450)  # hard coded
        index2 = randint(1, 1412)   # 7*7=49. 1456-50=1412
        # there will be 6 values
        data1,label1 = get_svm_index_based_method1(index1,time_list,value_list)
        data2,label2 = get_svm_index_based_method2(index2,time_list,value_list)
        # and the same 6 labels

        linear_svr = SVR(kernel='linear')
        linear_svr.fit(data1, label1)
        to_predict = np.array(['13']).reshape(-1, 1)
        print("test result is {}, prediction result is {} ".format(
            value_list[index1+7], linear_svr.predict(to_predict)[0]))
        expected_result_list1.append(value_list[index1+7])
        predict_result_list1.append(linear_svr.predict(to_predict)[0])
        linear_svr.fit(data2,label2)
        expected_result_list2.append(label2[-1])
        predict_result_list2.append(linear_svr.predict(to_predict)[0])

    # draw the plot1 for method1 sampling
    plt.plot(expected_result_list1)
    plt.plot(predict_result_list1)
    plt.ylim(ymin=0)

    plt.show()
    plt.clf()

    # draw the plot2 for method2 sampling
    plt.plot(expected_result_list2)
    plt.plot(predict_result_list2)
    plt.ylim(ymin=0)
    plt.show()


def get_neural_based_method1(index,time_list,value_list):
    data = [[1, 1],
            [2, 2],
            [3, 3],
            [4, 4],
            [5, 5],
            [6, 6],
            [7, 7]]
    label = np.array(value_list[index:index+7])
    return (data, label)

def get_neural_based_method2(index,time_list,value_list):
    data = [[1, 1],
            [2, 2],
            [3, 3],
            [4, 4],
            [5, 5],
            [6, 6],
            [7, 7]]
    label = np.array(list([value_list[index+7*i] for i in range(7)]))
    return data,label

def neural_network_predict():
    """ analyze by neural network
    """

    expected_result_list1 = []
    predict_result_list1 = []
    expected_result_list2 = []
    predict_result_list2 = []

   
    from sklearn.preprocessing import StandardScaler
    from sklearn.neural_network import MLPRegressor
    scaler = StandardScaler()


    for _ in range(200):
        index1 = randint(1, 1450)  # hard coded
        index2 = randint(1, 1412)   # 7*7=49. 1456-50=1412
        data1,label1 = get_neural_based_method1(index1,time_list,value_list)
        data2,label2 = get_neural_based_method2(index2,time_list,value_list)
    # data = np.array(time_list).reshape(-1, 1)  # .astype(np.float64)
    # label = np.array(value_list[:-1])  # .astype(np.float64)  # .reshape(-1, 1)

        mlp = MLPRegressor(solver='lbfgs', alpha=1e-5,
                        hidden_layer_sizes=(5, 2), random_state=1)
        mlp.fit(data1[:-1], label1[:-1])
        predictions = mlp.predict([[7, 7]])
        expected_result_list1.append(label1[-1])
        predict_result_list1.append(predictions[0] if predictions[0]>0 else 0)
        mlp.fit(data2[:-1],label2[:-1])
        expected_result_list2.append(label2[-1])
        predictions = mlp.predict([[7, 7]])
        predict_result_list2.append(predictions[0] if predictions[0]>0 and predictions[0]<50000 else 0)
    print(predict_result_list1)
    plt.plot(expected_result_list1)
    plt.plot(predict_result_list1)
    plt.ylim(ymin=0)

    plt.show()
    plt.clf()

    # draw the plot2 for method2 sampling
    plt.plot(expected_result_list2)
    plt.plot(predict_result_list2)
    plt.ylim(ymin=0)
    plt.show()


def time_series_analysis():

    from pandas import read_csv
    from pandas import datetime
    from matplotlib import pyplot

    def parser(x):
        return datetime.strptime('190'+x, '%Y-%m')

    series = read_csv('elec.csv', header=0,
                      parse_dates=[0], index_col=0, squeeze=True, date_parser=parser)
    series = {  # 'time': pd.Series(['2011-01-01', '2011-01-02', '2011-01-03', '2011-01-04', '2011-01-05', '2011-01-06']),
        'time': pd.Series(['1901-01', '1902-01', '1903-01', '1904-01', '1905-01', '1906-01', '1907-01', '1908-01']),
        'value': pd.Series([19330.143540669856600, 30641.148325358849700, 23813.397129186604700, 23272.727272727275100,
                            22866.028708133973200, 23961.722488038278900, 25856.459330143542400, 29598.086124401913600])}
    series = pd.DataFrame(series)
    print(series.head())
#     series.plot()
#     pyplot.show()
    from sklearn.metrics import mean_squared_error
    from pandas import read_csv
    from statsmodels.tsa.arima_model import ARIMA
    series = read_csv('shampoo.csv', header=0, parse_dates=[
                      0], index_col=0, squeeze=True)
    # fit model
    model = ARIMA(series, order=(5, 1, 0))
    model_fit = model.fit(disp=0)
    print(model_fit.summary())
    # plot residual errors
    residuals = DataFrame(model_fit.resid)
    residuals.plot()
#     pyplot.show()
    residuals.plot(kind='kde')
#     pyplot.show()
    print(residuals.describe())
    X = series.values
    size = int(len(X) * 0.66)
    train, test = X[0:size], X[size:len(X)]
    history = [x for x in train]
    predictions = list()
    for t in range(len(test)):
        model = ARIMA(history, order=(5, 1, 0))
        model_fit = model.fit(disp=0)
        output = model_fit.forecast()
        yhat = output[0]
        predictions.append(yhat)
        obs = test[t]
        history.append(obs)
        print('predicted=%f, expected=%f' % (yhat, obs))
    error = mean_squared_error(test, predictions)
    print('Test MSE: %.3f' % error)
    print(predictions)



if __name__ == '__main__':
    # svm_predict()
    # neural_network_predict()
    time_series_analysis()
#     test_series_analysis()
