import matplotlib.pyplot as plt
times = []
values = []
sum = 0
index = 0
with open("MT123electricity.csv") as input_:
    for line in input_:
        index += 1
        times.append(line.split(",")[0])
        values.append(line.split(",")[1])
        sum += float(line.split(",")[1][1:-2].strip("\n"))
        if(index > 1):
            break

print(sum)
# plt.plot(times, values)
# plt.show()


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


if __name__ == '__main__':
    svm_predict()
