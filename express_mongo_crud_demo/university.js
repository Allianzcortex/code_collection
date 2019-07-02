var SERVER_URL = "http://dev.cs.smu.ca:8171";
//var SERVER_URL = "http://127.0.0.1:8171"


function saveUniversity() {
  // $("#universityForm").validate()
  if ($("#name").val() == "" || $("#phone").val() == "" || $("#address").val() == "") {
    $("#name").focus();
    $("#error_alert").append("All three options cannot be null");
    setTimeout(function () {
      $("#error_alert").empty();
    }, 800);
  } else {

    var param = {
      Name: $("#name").val(),
      Address: $("#address").val(),
      PhoneNumber: $("#phone").val()
    }
    console.log(SERVER_URL + "/saveUniversity");
    $.post(SERVER_URL + "/saveUniversity", param, function (data) {
      alert(data);
    }).fail(function (error) {
      if (error.responseText == "") {
        alert("There is no available running server");
      } else {
        console.log(error);
        alert(error);
      }
    });
  }
}

  function deleteUniversity() {
    if ($("#name").val() == "") {
      $("#name").focus();
      $("#error_alert").append("delete name cannot be null");
      setTimeout(function () {
        $("#error_alert").empty();
      }, 800);

    } else {
      var param = {
        Name: $("#name").val(),
        Address: $("#address").val(),
        PhoneNumber: $("#phone").val()
      };
      $.post(SERVER_URL + "/deleteUniversity", param, function (data) {
        alert(data);
      }).fail(function (error) {
        if (error.responseText == "") {
          alert("There is no available running server");
        } else {
          console.log(error);
          alert(error.responseText);
        }
      });
    }
  }

function searchRecord() {
  if ($("#search_keyword").val() == "") {
    $("#search_keyword").focus();
    $("#error_alert").append("search key word cannot be null");
    setTimeout(function () {
      $("#error_alert").empty();
    }, 2000);
  } else {
    var newObj = {
      Name: $("#search_keyword").val()
    };

    $.post(SERVER_URL + "/searchUniversity", newObj, function (data) {
      var trHTML = 'Below are all the records:' + '<thead>\n' +
        '    <tr>\n' +
        '      <th scope="col">University Name</th>\n' +
        '      <th scope="col">Email Address</th>\n' +
        '      <th scope="col">Phone Number</th>\n' +
        '    </tr>\n' +
        '  </thead>\n' +
        '            <tbody>';

      var result = [];
      for (i = 0; i < data.length; i++) {
        result.push([])
        result[i].push(data[i]["Name"])
        result[i].push(data[i]["Address"])
        result[i].push(data[i]["PhoneNumber"])
      }

      var $select = $('#search_result');
      var listitems = '';
      for (i = 0; i < result.length; i++) {
        trHTML += '<tr>';
        for (j = 0; j < result[i].length; j++) {
          trHTML += '<td>' + result[i][j] + '</td>';
        }
        trHTML += '/<tr>';
        trHTML += '/<tbody>';
      }
      $('#search_result').empty();
      $('#search_result').append(trHTML);

    }).fail(function (error) {
      if (error.responseText == "") {
        alert("There is no available running server");
      } else {
        console.log(error);
        alert(error.responseText);
      }
    });
  }
}



function displayRecords() {
  $.post(SERVER_URL + "/displayUniversity", function (data) {
    var trHTML = 'Below are all the records:' + '<thead>\n' +
      '    <tr>\n' +
      '      <th scope="col">University Name</th>\n' +
      '      <th scope="col">Email Address</th>\n' +
      '      <th scope="col">Phone Number</th>\n' +
      '    </tr>\n' +
      '  </thead>\n' +
      '            <tbody>';

    var result = [];
    for (i = 0; i < data.length; i++) {
      result.push([])
      result[i].push(data[i]["Name"])
      result[i].push(data[i]["Address"])
      result[i].push(data[i]["PhoneNumber"])
    }

    var $select = $('#search_result');
    var listitems = '';
    for (i = 0; i < result.length; i++) {
      trHTML += '<tr>';
      for (j = 0; j < result[i].length; j++) {
        trHTML += '<td>' + result[i][j] + '</td>';
      }
      trHTML += '/<tr>';
      trHTML += '/<tbody>';
    }
    $('#search_result').empty();
    $('#search_result').append(trHTML);
  }).fail(function (error) {
    if (error.responseText == "") {
      alert("There is no available running server");
    } else {
      console.log(error);
      alert(error.responseText);
    }
  });
}