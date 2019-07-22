from pandas import datetime
from matplotlib import pyplot
from statsmodels.tsa.arima_model import ARIMA
from pandas import DataFrame
import matplotlib.pyplot as plt
import numpy as np
times = []
values = []
sum = 0
index = 0

date = ['01',
        '02',
        '03',
        '04',
        '05',
        '06',
        '07',
        '08',
        '09',
        '10',
        '11',
        '12',
        '13'
        ]

date = [1,
        2,
        3,
        4,
        5,
        6,
        7,
        8,
        9,
        10,
        11,
        12,
        13
        ]
value = [19330.143540669856600,
         30641.148325358849700,
         23813.397129186604700,
         23272.727272727275100,
         22866.028708133973200,
         23961.722488038278900,
         25856.459330143542400,
         29598.086124401913600,
         30263.157894736844900,
         24167.464114832536600,
         23047.846889952151200,
         23995.215311004792000,
         23679.425837320575800]

value = [1,
         3,
         2,
         2,
         2,
         2,
         2,
         2,
         3,
         2,
         2,
         2,
         2]


with open("MT123electricity.csv") as input_:
    for line in input_:
        index += 1
        times.append(line.split(",")[0])
        values.append(line.split(",")[1])
        sum += float(line.split(",")[1][1:-2].strip("\n"))
        if(index > 1):
            break

# print(sum)
# plt.plot(times, values)
# plt.show()


def neural_network_predict():
    """ analyze by neural network
    """
    from sklearn.preprocessing import StandardScaler
    scaler = StandardScaler()
    # X_train =
    # scaler.fit(X_train)
    # X_train = scaler.transform(X_train)
    # X_test = scaler.transform(X_test)
    # global value
    # data = date[:-1]
    # print(data[-1])
    # data = np.array([data]).T
    # value = [int(x) for x in value]
    # label = np.array(value[:-1])
    global date
    global value
    value = list(map(int, value))
    print(type(value[0]))
    print(value)
    date = np.array(date[:-1]).reshape(-1, 1)  # .astype(np.float64)
    value = np.array(value[:-1])  # .astype(np.float64)  # .reshape(-1, 1)
    from sklearn.neural_network import MLPRegressor
    mlp = MLPRegressor(solver='lbfgs', alpha=1e-5,
                       hidden_layer_sizes=(5, 2), random_state=1)
    print(date)
    print(type(value[0]))
    print(type(date[0]))

    date = [[1, 1],
            [2, 2],
            [3, 3],
            [4, 4],
            [5, 5],
            [6, 6],
            [7, 7],
            [8, 8],
            [9, 9],
            [10, 10],
            [11, 11],
            [12, 12],
            [13, 13]]
   # date = np.array(date).reshape(1, -1)
    value = [1,
             3,
             2,
             2,
             2,
             2,
             2,
             2,
             3,
             2,
             2,
             2,
             2]
    value = [19330.143540669856600,
             30641.148325358849700,
             23813.397129186604700,
             23272.727272727275100,
             22866.028708133973200,
             23961.722488038278900,
             25856.459330143542400,
             29598.086124401913600,
             30263.157894736844900,
             24167.464114832536600,
             23047.846889952151200,
             23995.215311004792000,
             23679.425837320575800]
    # print(date.shape)
    # value=np.array(value).reshape(1,-1)
    mlp.fit(date, value)
    # X = [[0., 0.], [1., 1.]]
    # y = [0, 1]
    # mlp.fit(X, y)
    # print(mlp.predict([[2., 2.]]))
    predictions = mlp.predict([[13, 13]])
    print(predictions)


def svm_predict():
    # date = ['2011-01-01',
    #         '2011-01-02',
    #         '2011-01-03',
    #         '2011-01-04',
    #         '2011-01-05',
    #         '2011-01-06',
    #         '2011-01-07',
    #         '2011-01-08',
    #         '2011-01-09',
    #         '2011-01-10',
    #         '2011-01-11',
    #         '2011-01-12',
    #         '2011-01-13'
    #         ]
    date = [1,
            2,
            3,
            4,
            5,
            6,
            7,
            8,
            9,
            10,
            11,
            12,
            13
            ]
    value = [19330.143540669856600,
             30641.148325358849700,
             23813.397129186604700,
             23272.727272727275100,
             22866.028708133973200,
             23961.722488038278900,
             25856.459330143542400,
             29598.086124401913600,
             30263.157894736844900,
             24167.464114832536600,
             23047.846889952151200,
             23995.215311004792000,
             23679.425837320575800]

    from sklearn.svm import SVR
    import numpy as np
    data = date[:-1]
    print(data[-1])
    data = np.array([data]).T
    label = np.array(value[:-1])
    linear_svr = SVR(kernel='linear')
    linear_svr.fit(data, label)
    to_predict = np.array(['13']).reshape(-1, 1)
    print("test result is {}, prediction result is {} ".format(
        23679.425837320575800, linear_svr.predict(to_predict)))


def test_series_analysis():

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


def time_series_analysis():
    import numpy as np
    import pandas as pd
#   data = pd.read_csv("Electric_Production.csv", index_col=0)

    data = {  # 'time': pd.Series(['2011-01-01', '2011-01-02', '2011-01-03', '2011-01-04', '2011-01-05', '2011-01-06']),
        'time': pd.Series(['1.0', '2.0', '3.0', '4.0', '5.0', '6.0', '7.0', '8.0']),
        'value': pd.Series([19330.143540669856600, 30641.148325358849700, 23813.397129186604700, 23272.727272727275100,
                            22866.028708133973200, 23961.722488038278900, 25856.459330143542400, 29598.086124401913600])}
    data = pd.DataFrame(data)
    data = pd.read_csv("elec.csv")

    data.index = pd.to_datetime(data.index)
#     from pyramid.arima import auto_arima
    from statsmodels.tsa.arima_model import ARIMA
    model = ARIMA(data.astype(float), order=(3, 1, 2))
    model_fit = model.fit(disp=0)
    print(model_fit.forecast()[0])
#     stepwise_model = auto_arima(data, start_p=1, start_q=1,
#                                 max_p=3, max_q=3, m=12,
#                                 start_P=0, seasonal=True,
#                                 d=1, D=1, trace=True,
#                                 error_action='ignore',
#                                 suppress_warnings=True,
#                                 stepwise=True)
#     stepwise_model.fit(train)
#     future_forecast = stepwise_model.predict(n_periods=37)
#     future_forecast


if __name__ == '__main__':
    svm_predict()
    #     neural_network_predict()
    #     time_series_analysis()
#     test_series_analysis()
