function myFunction() {
    $("#records_table tr").remove();
    // var tableName = document.getElementById("down").value;
    var tableName = $("#down").val();

    $.ajax({
        type: "GET",
        url: "http://127.0.0.1:8080/api/getTable/" + tableName,
        // data:{Name:"sanmao",Password:"sanmaoword"},
        headers: {
            'Access-Control-Allow-Origin': '*'
        },
        dataType: 'json',

        success: function (data) {

            // var obj = jQuery.parseJSON(data);
            console.log(data.values);
            var ffhtml = '<tr>';
            data.fields.forEach(function (entry) {
                ffhtml = ffhtml + '<th>' + entry + '</th>';
            });
            ffhtml += '</tr>';
            $('#records_table').append(ffhtml);
            var trHTML = '';
            $.each(data.values, function (i, item) {
                trHTML += '<tr>';
                // https://javascript.info/keys-values-entries
                // https://stackoverflow.com/questions/9329446/for-each-over-an-array-in-javascript
                var s = Object.values(item);

                for (i = 0; i < s.length; i++) {
                    trHTML += '<td>' + s[i] + '</td>';
                }

                trHTML += '/<tr>';
                // trHTML += '<tr><td>' + item.id + '</td><td>' + item.price  + '</td></tr>';
            });
            $('#records_table').append(trHTML);
        },

        complete: function (XMLHttpRequest, textStatus) {
            // alert(XMLHttpRequest.responseText);
            // alert(textStatus);
            //HideLoading();
        },
        //调用出错执行的函数
        error: function () {

        }
    });
}


$(document).ready(function () {
    // click on button submit
    $("#submit").on('click', function () {
        var str = document.getElementById("authors").value;

        var tempResult = str.split(";");
        var authorsList = tempResult.map(function (x) {
            return parseInt(x, 10);
        });

        var article = {
            magazineId: document.getElementById("magazine").value,
            volume: document.getElementById("volume").value,
            volumeNumber: parseInt(document.getElementById("volumenumber").value),
            title: document.getElementById("title").value,
            pages: parseInt(document.getElementById("pages").value),
            publicationYear: document.getElementById("publication_year").value,
            author: document.getElementById("authors").value
        }
        console.log(document.getElementById("authors").value);
        console.log(article);
        // send ajax
        $.ajax({
            url: 'http://localhost:8080/api/create/article', // url where to submit the request
            type: "POST", // type of action POST || GET
            dataType: 'json', // data type
            contentType: "application/json",
            // data : $("#form").serialize(), // post data || get data
            data: JSON.stringify(article),
            success: function (result) {
                // you can see the result from the console
                // tab of the developer tools

               if (parseInt(result) === -1) {
                    $("#error_1").text('The volume number cannot be repeated').css('color', 'red');
                } else if (parseInt(result) === -2) {
                    $("#error_2").text('The same magazine volume should have the same publication year').css('color', 'red');
                } else if (parseInt(result) === -3) {
                    $("#error_5").text('The name of magazine cannot be empty').css('color', 'red');
                } else {
                    alert("Add Article Successfully!")
                }
            },
            error: function (xhr, resp, text) {
                console.log(xhr, resp, text);
            }
        })
    });

    // create new customer
    $("#submit_new_customer").on('click', function () {
        var customer = {
            firstName: document.getElementById("firstName").value,
            lastName: document.getElementById("lastName").value,
            telephoneNumber: document.getElementById("telephoneNumber").value,
            mailingAddress: document.getElementById("mailingAddress").value

        }

        $.ajax({
            url: 'http://localhost:8080/api/create/customer?checkExists=true', // url where to submit the request
            type: "POST", // type of action POST || GET
            dataType: 'json', // data type
            contentType: "application/json",
            // data : $("#form").serialize(), // post data || get data
            data: JSON.stringify(customer),
            success: function (result) {
                // you can see the result from the console
                // tab of the developer tools
                console.log(result);
                alert("Add new customer successfully");
            },
            error: function (xhr, resp, text) {

                var r = confirm("The customer already exists in database,Are you sure this is a new cusotmer?");
                if (r === true) {
                    $.ajax({
                        url: 'http://localhost:8080/api/create/customer?checkExists=false', // url where to submit the request
                        type: "POST", // type of action POST || GET
                        dataType: 'json', // data type
                        contentType: "application/json",
                        // data : $("#form").serialize(), // post data || get data
                        data: JSON.stringify(customer)
                    });
                    txt = "A new customer inserted";
                } else {
                    txt = "Add operation cancelled!";
                }


            }
        })
    });


    // create new transaction
    $("#submit_new_transaction").on('click', function () {
        var str = document.getElementById("items").value;
        // 1:12;2:13
        var resultArray = [];
        var tempResult = str.split(";");
        tempResult.map(function (x) {
            var tempResult = {
                "id": x.split(":")[0],
                "price": parseInt(x.split(":")[1])
            }
            resultArray.push(tempResult);

        });
        console.log("----");
        console.log(resultArray);

        var transaction = {
            customerId: document.getElementById("customerId").value,
            items: resultArray

        }

        $.ajax({
            url: 'http://localhost:8080/api/create/transaction', // url where to submit the request
            type: "POST", // type of action POST || GET
            dataType: 'json', // data type
            contentType: "application/json",
            // data : $("#form").serialize(), // post data || get data
            data: JSON.stringify(transaction),
            success: function (result) {
                // you can see the result from the console
                // tab of the developer tools
                alert("Add new transaction successfully");

            },
            error: function (xhr, resp, text) {
                $("#error_8").text('The item does not exist').css('color', 'red');
                console.log(xhr, resp, text);
            }
        })

    });

    // delete transaction
    $("#delete_transaction").on('click', function () {

        $.ajax({
            url: 'http://localhost:8080/api/cancel/transaction/' + parseInt(document.getElementById("transactionId").value), // url where to submit the request
            type: "DELETE", // type of action POST || GET
            dataType: 'json', // data type
            // data : $("#form").serialize(), // post data || get data
            success: function (result) {
                // you can see the result from the console
                // tab of the developer tools
                console.log(result);
                alert("Delete transaction successfully");
                var trHTML = '<thead>\n' +
                    '        <tr>\n' +
                    '            <th scope="col">transactionNumber</th>\n' +
                    '            <th scope="col">transactionDate</th>\n' +
                    '            <th scope="col">totalPurchasePrice</th>\n' +
                    '            <th scope="col">customerId</th>\n' +
                    '        </tr>\n' +
                    '        </thead>';

                $.each(result, function (i, item) {
                    trHTML += '<tr><td>' + item.transactionNumber + '</td><td>' + item.transactionDate + '</td><td>' + item.totalPurchasePrice + '</td><td>' +
                        item.customerId + '</td></tr>';
                });
                trHTML += '</tbody>'
                $('#records_table').append(trHTML);
            },
            error: function (xhr, resp, text) {
                var s = JSON.stringify(xhr.responseText);
                var i = s.indexOf("message");
                var errorMsg = s.substring(i + 12, i + 14);
                if (errorMsg === "-1") {
                    $("#error_4").text('The Transaction Number doesn\'t exist').css('color', 'red');
                } else if (errorMsg === "-2") {
                    $("#error_4").text('The Transaction  happened before 1 month , it cann\'t be deleted').css('color', 'red');
                }
            }
            // error: function (err) {
            //     alert(JSON.parse(err))
            //     alert(JSON.stringify(err));
            // }
        })


    });

});