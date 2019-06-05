from sklearn.preprocessing import LabelEncoder,LabelBinarizer
from sklearn.tree import DecisionTreeClassifier 
from sklearn.model_selection import train_test_split
from sklearn.metrics import accuracy_score
from sklearn.metrics import confusion_matrix
from sklearn.tree import export_graphviz
from sklearn.externals.six import StringIO
from sklearn.metrics import roc_curve, roc_auc_score,confusion_matrix,auc,recall_score
from sklearn.multiclass import OneVsRestClassifier

import matplotlib.pyplot as plt
import pandas as pd
from IPython.display import Image
import pydotplus


def read_data():

    pima = pd.read_csv("car.data.txt", header=0)
    print(pima.head()) 

    le = LabelEncoder()
    categories_non_numerical = [
        'price', 'maintenance', 'storage', 'safety', 'shouldBuy']
    pima[categories_non_numerical] = pima[categories_non_numerical].apply(
        lambda col: le.fit_transform(col))
    print(set(pima['price']))

    pima['doors'].replace('5more', '5', inplace=True)
    pima['seats'].replace('more', '5', inplace=True)

    # split dataset in features and target variable
    feature_cols = ['price', 'maintenance',
                    'doors', 'seats', 'storage', 'safety']
    X = pima[feature_cols]  # Features

    target_cols = ['shouldBuy']
    y = pima[target_cols]  # Target variable

    X_train, X_test, y_train, y_test = train_test_split(
        X, y, test_size=0.3, random_state=1)  # 70% training and 30% test

    return X_train, X_test, y_train, y_test


def generate_decision_tree_classifier():

    # Create Decision Tree classifer object
    # e.g.：max_depth, max_leaf_nodes, min_impurity_split, and min_samples_leaf
    clf = DecisionTreeClassifier()

    X_train, X_test, y_train, y_test = read_data()
    # Train Decision Tree Classifer
    clf = clf.fit(X_train, y_train)

    # Predict the response for test dataset
    y_predict = clf.predict(X_test)

    print("below is the predict result and test result")
    print(y_predict)
    print(y_test)

    # about K-fold test
    from sklearn.model_selection import cross_val_score
    # kf = KFold(25, n_folds=5, shuffle=False)
    depth = []
    for i in range(1,50):
        clf = DecisionTreeClassifier(max_depth=i)
        # Perform 7-fold cross validation 
        scores = cross_val_score(estimator=clf, X=X_train, y=y_train, cv=7, n_jobs=4)
        depth.append((i,scores.mean()))
    print(depth)
    x=[f[0] for f in depth]
    y=[f[1] for f in depth]
    print("x is ")
    print(x)
    print(y)
    plt.scatter(x,y)
    plt.title("relationship between")
    plt.xlabel("max_depth")
    plt.ylabel("accuracy")
    plt.show()

    # Model Accuracy, how often is the classifier correct?
    print("Accuracy:", accuracy_score(y_test, y_predict))
    print("recall_score: ",recall_score(y_test,y_predict, average='macro'))
    
    # calculate confusion matrix
    # Just For random test
    cm = pd.DataFrame(
        confusion_matrix(y_test, y_predict),
        columns=['l', '2', '3', '4'],
        index=['1', '2', '3', '4']
    )
    print(cm)

    # auc and roc
    lb = LabelBinarizer()
    lb.fit(y_test)
    y_test1 = lb.transform(y_test)
    y_pred = lb.transform(y_predict)
    print(roc_auc_score(y_test1, y_pred, average='macro'))

    n_classes = 4
    """
    TODO: cannot find a better way to deal with multiple-class
    roc and auc plot.
    So convert it into two-classes problem
    """
    #  predict
    print(type(y_predict))
    y_predict[y_predict <= 1]=0
    y_predict[y_predict >= 2]=1
    #  y_test
    y_test.replace(1,0,inplace=True)
    y_test.replace(2,1,inplace=True)
    y_test.replace(3,1,inplace=True)
 
    fpr, tpr, thresholds = roc_curve(y_test,y_predict)
    plt.plot([0, 1], [0, 1], linestyle='--')
    plt.plot(fpr, tpr, marker='.')
    plt.show()
    
    # clf = OneVsRestClassifier(DecisionTreeClassifier(random_state=0))

    # X_train, X_test, y_train, y_test = read_data()
    # # Train Decision Tree Classifer
    # y_score = clf.fit(X_train, y_train).predict_proba(X_test)
    # fpr = dict()
    # tpr = dict()
    # roc_auc = dict()
    # for i in range(n_classes):
    #     fpr[i], tpr[i], _ = roc_curve(y_test[:, i], y_score[:, i])
    #     roc_auc[i] = auc(fpr[i], tpr[i])
    # colors = cycle(['blue', 'red', 'green'])
    # for i, color in zip(range(n_classes), colors):
    #     plt.plot(fpr[i], tpr[i], color=color, lw=lw,
    #              label='ROC curve of class {0} (area = {1:0.2f})'
    #              ''.format(i, roc_auc[i]))


def generate_random_forest_classfier():
    # 500 means the number of ntrees
    from sklearn.ensemble import RandomForestClassifier
    clf = RandomForestClassifier(n_estimators=600, criterion="entropy")
    X_train, X_test, y_train, y_test = read_data()
    clf.fit(X_train, y_train)

    print(clf)
    predicitions = clf.predict(X_test)

    print("Accuracy:", accuracy_score(y_test, predicitions))
    feature_importances = pd.DataFrame(clf.feature_importances_,
                                       index=X_train.columns,
                                       columns=['importance']) .sort_values('importance',ascending = False)
    importance = list(feature_importances.iloc[:,0])
    names = list(feature_importances.index)
   
    plt.bar(names,importance)
    plt.show()



if __name__ == '__main__':
    # uncomment the following comment to run 2 classifer trees both
    #generate_random_forest_classfier()
    generate_decision_tree_classifier()

    # graph.write_png('diabetes.png')
    # Image(graph.create_png())
