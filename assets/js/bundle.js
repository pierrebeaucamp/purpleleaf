(function e(t,n,r){function s(o,u){if(!n[o]){if(!t[o]){var a=typeof require=="function"&&require;if(!u&&a)return a(o,!0);if(i)return i(o,!0);var f=new Error("Cannot find module '"+o+"'");throw f.code="MODULE_NOT_FOUND",f}var l=n[o]={exports:{}};t[o][0].call(l.exports,function(e){var n=t[o][1][e];return s(n?n:e)},l,l.exports,e,t,n,r)}return n[o].exports}var i=typeof require=="function"&&require;for(var o=0;o<r.length;o++)s(r[o]);return s})({1:[function(require,module,exports){
"use strict";

$(document).ready(function () {
  function PaymillResponseHandler(error, result) {
    if (error) {
      // Displays the error above the form
      $(".payment-errors").text(error.apierror);
    } else {
      $(".payment-errors").text("");
      var form = $("#payment-form");
      // Token
      var token = result.token;

      // Insert token into form in order to submit to server
      form.append("<input type='hidden' name='paymillToken' value='" + token + "'/>");
      form.get(0).submit();
    }
    $(".submit-button").removeAttr("disabled");
  }

  $("#payment-form").on("submit", function (event) {
    // Deactivate submit button to avoid further clicks
    $(".submit-button").attr("disabled", "disabled");

    if (!paymill.validateCardNumber($(".card-number").val())) {
      $(".payment-errors").text("Invalid card number");
      $(".submit-button").removeAttr("disabled");
      return false;
    }

    if (!paymill.validateExpiry($(".card-expiry-month").val(), $(".card-expiry-year").val())) {
      $(".payment-errors").text("Invalid expiration date");
      $(".submit-button").removeAttr("disabled");
      return false;
    }

    console.log(paymill);

    paymill.createToken({
      number: $(".card-number").val(),
      exp_month: $(".card-expiry-month").val(),
      exp_year: $(".card-expiry-year").val(),
      cvc: $(".card-cvc").val(),
      cardholder: $(".card-holdername").val(),
      amount_int: $(".card-amount-int").val(), // Integer e.g. "4900" fÃ¼r 49,00 EUR
      currency: $(".card-currency").val() // ISO 4217 z.B. "EUR"
    }, PaymillResponseHandler);

    return false;
  });
});

},{}]},{},[1])
//# sourceMappingURL=data:application/json;charset:utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIm5vZGVfbW9kdWxlcy9icm93c2VyaWZ5L25vZGVfbW9kdWxlcy9icm93c2VyLXBhY2svX3ByZWx1ZGUuanMiLCIvVXNlcnMvcGF0cmlja2dvbGxtZXIvU2l0ZXMvcHVycGxlbGVhZi9zcmMvamF2YXNjcmlwdC9wYWdlLmpzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiJBQUFBOzs7QUNBQSxDQUFDLENBQUMsUUFBUSxDQUFDLENBQUMsS0FBSyxDQUFDLFlBQVc7QUFDM0IsV0FBUyxzQkFBc0IsQ0FBQyxLQUFLLEVBQUUsTUFBTSxFQUFFO0FBQzdDLFFBQUksS0FBSyxFQUFFOztBQUVULE9BQUMsQ0FBQyxpQkFBaUIsQ0FBQyxDQUFDLElBQUksQ0FBQyxLQUFLLENBQUMsUUFBUSxDQUFDLENBQUM7S0FDM0MsTUFBTTtBQUNMLE9BQUMsQ0FBQyxpQkFBaUIsQ0FBQyxDQUFDLElBQUksQ0FBQyxFQUFFLENBQUMsQ0FBQztBQUM5QixVQUFJLElBQUksR0FBRyxDQUFDLENBQUMsZUFBZSxDQUFDLENBQUM7O0FBRTlCLFVBQUksS0FBSyxHQUFHLE1BQU0sQ0FBQyxLQUFLLENBQUM7OztBQUd6QixVQUFJLENBQUMsTUFBTSxDQUFDLGtEQUFrRCxHQUFHLEtBQUssR0FBRyxLQUFLLENBQUMsQ0FBQztBQUNoRixVQUFJLENBQUMsR0FBRyxDQUFDLENBQUMsQ0FBQyxDQUFDLE1BQU0sRUFBRSxDQUFDO0tBQ3RCO0FBQ0QsS0FBQyxDQUFDLGdCQUFnQixDQUFDLENBQUMsVUFBVSxDQUFDLFVBQVUsQ0FBQyxDQUFDO0dBQzVDOztBQUVELEdBQUMsQ0FBQyxlQUFlLENBQUMsQ0FBQyxFQUFFLENBQUMsUUFBUSxFQUFFLFVBQVMsS0FBSyxFQUFFOztBQUU5QyxLQUFDLENBQUMsZ0JBQWdCLENBQUMsQ0FBQyxJQUFJLENBQUMsVUFBVSxFQUFFLFVBQVUsQ0FBQyxDQUFDOztBQUVqRCxRQUFJLENBQUMsT0FBTyxDQUFDLGtCQUFrQixDQUFDLENBQUMsQ0FBQyxjQUFjLENBQUMsQ0FBQyxHQUFHLEVBQUUsQ0FBQyxFQUFFO0FBQ3hELE9BQUMsQ0FBQyxpQkFBaUIsQ0FBQyxDQUFDLElBQUksQ0FBQyxxQkFBcUIsQ0FBQyxDQUFDO0FBQ2pELE9BQUMsQ0FBQyxnQkFBZ0IsQ0FBQyxDQUFDLFVBQVUsQ0FBQyxVQUFVLENBQUMsQ0FBQztBQUMzQyxhQUFPLEtBQUssQ0FBQztLQUNkOztBQUVELFFBQUksQ0FBQyxPQUFPLENBQUMsY0FBYyxDQUN6QixDQUFDLENBQUMsb0JBQW9CLENBQUMsQ0FBQyxHQUFHLEVBQUUsRUFDN0IsQ0FBQyxDQUFDLG1CQUFtQixDQUFDLENBQUMsR0FBRyxFQUFFLENBQUMsRUFDN0I7QUFDQSxPQUFDLENBQUMsaUJBQWlCLENBQUMsQ0FBQyxJQUFJLENBQUMseUJBQXlCLENBQUMsQ0FBQztBQUNyRCxPQUFDLENBQUMsZ0JBQWdCLENBQUMsQ0FBQyxVQUFVLENBQUMsVUFBVSxDQUFDLENBQUM7QUFDM0MsYUFBTyxLQUFLLENBQUM7S0FDZDs7QUFFRCxXQUFPLENBQUMsR0FBRyxDQUFDLE9BQU8sQ0FBQyxDQUFBOztBQUVwQixXQUFPLENBQUMsV0FBVyxDQUFDO0FBQ2xCLFlBQU0sRUFBVSxDQUFDLENBQUMsY0FBYyxDQUFDLENBQUMsR0FBRyxFQUFFO0FBQ3ZDLGVBQVMsRUFBTyxDQUFDLENBQUMsb0JBQW9CLENBQUMsQ0FBQyxHQUFHLEVBQUU7QUFDN0MsY0FBUSxFQUFRLENBQUMsQ0FBQyxtQkFBbUIsQ0FBQyxDQUFDLEdBQUcsRUFBRTtBQUM1QyxTQUFHLEVBQWEsQ0FBQyxDQUFDLFdBQVcsQ0FBQyxDQUFDLEdBQUcsRUFBRTtBQUNwQyxnQkFBVSxFQUFNLENBQUMsQ0FBQyxrQkFBa0IsQ0FBQyxDQUFDLEdBQUcsRUFBRTtBQUMzQyxnQkFBVSxFQUFNLENBQUMsQ0FBQyxrQkFBa0IsQ0FBQyxDQUFDLEdBQUcsRUFBRTtBQUMzQyxjQUFRLEVBQVEsQ0FBQyxDQUFDLGdCQUFnQixDQUFDLENBQUMsR0FBRyxFQUFFO0FBQUEsS0FDeEMsRUFBRSxzQkFBc0IsQ0FBQyxDQUFDOztBQUU3QixXQUFPLEtBQUssQ0FBQztHQUNkLENBQUMsQ0FBQztDQUNKLENBQUMsQ0FBQyIsImZpbGUiOiJnZW5lcmF0ZWQuanMiLCJzb3VyY2VSb290IjoiIiwic291cmNlc0NvbnRlbnQiOlsiKGZ1bmN0aW9uIGUodCxuLHIpe2Z1bmN0aW9uIHMobyx1KXtpZighbltvXSl7aWYoIXRbb10pe3ZhciBhPXR5cGVvZiByZXF1aXJlPT1cImZ1bmN0aW9uXCImJnJlcXVpcmU7aWYoIXUmJmEpcmV0dXJuIGEobywhMCk7aWYoaSlyZXR1cm4gaShvLCEwKTt2YXIgZj1uZXcgRXJyb3IoXCJDYW5ub3QgZmluZCBtb2R1bGUgJ1wiK28rXCInXCIpO3Rocm93IGYuY29kZT1cIk1PRFVMRV9OT1RfRk9VTkRcIixmfXZhciBsPW5bb109e2V4cG9ydHM6e319O3Rbb11bMF0uY2FsbChsLmV4cG9ydHMsZnVuY3Rpb24oZSl7dmFyIG49dFtvXVsxXVtlXTtyZXR1cm4gcyhuP246ZSl9LGwsbC5leHBvcnRzLGUsdCxuLHIpfXJldHVybiBuW29dLmV4cG9ydHN9dmFyIGk9dHlwZW9mIHJlcXVpcmU9PVwiZnVuY3Rpb25cIiYmcmVxdWlyZTtmb3IodmFyIG89MDtvPHIubGVuZ3RoO28rKylzKHJbb10pO3JldHVybiBzfSkiLCIkKGRvY3VtZW50KS5yZWFkeShmdW5jdGlvbigpIHtcbiAgZnVuY3Rpb24gUGF5bWlsbFJlc3BvbnNlSGFuZGxlcihlcnJvciwgcmVzdWx0KSB7XG4gICAgaWYgKGVycm9yKSB7XG4gICAgICAvLyBEaXNwbGF5cyB0aGUgZXJyb3IgYWJvdmUgdGhlIGZvcm1cbiAgICAgICQoXCIucGF5bWVudC1lcnJvcnNcIikudGV4dChlcnJvci5hcGllcnJvcik7XG4gICAgfSBlbHNlIHtcbiAgICAgICQoXCIucGF5bWVudC1lcnJvcnNcIikudGV4dChcIlwiKTtcbiAgICAgIHZhciBmb3JtID0gJChcIiNwYXltZW50LWZvcm1cIik7XG4gICAgICAvLyBUb2tlblxuICAgICAgdmFyIHRva2VuID0gcmVzdWx0LnRva2VuO1xuXG4gICAgICAvLyBJbnNlcnQgdG9rZW4gaW50byBmb3JtIGluIG9yZGVyIHRvIHN1Ym1pdCB0byBzZXJ2ZXJcbiAgICAgIGZvcm0uYXBwZW5kKFwiPGlucHV0IHR5cGU9J2hpZGRlbicgbmFtZT0ncGF5bWlsbFRva2VuJyB2YWx1ZT0nXCIgKyB0b2tlbiArIFwiJy8+XCIpO1xuICAgICAgZm9ybS5nZXQoMCkuc3VibWl0KCk7XG4gICAgfVxuICAgICQoXCIuc3VibWl0LWJ1dHRvblwiKS5yZW1vdmVBdHRyKFwiZGlzYWJsZWRcIik7XG4gIH1cbiAgXG4gICQoXCIjcGF5bWVudC1mb3JtXCIpLm9uKCdzdWJtaXQnLCBmdW5jdGlvbihldmVudCkge1xuICAgIC8vIERlYWN0aXZhdGUgc3VibWl0IGJ1dHRvbiB0byBhdm9pZCBmdXJ0aGVyIGNsaWNrc1xuICAgICQoJy5zdWJtaXQtYnV0dG9uJykuYXR0cihcImRpc2FibGVkXCIsIFwiZGlzYWJsZWRcIik7XG4gICAgXG4gICAgaWYgKCFwYXltaWxsLnZhbGlkYXRlQ2FyZE51bWJlcigkKCcuY2FyZC1udW1iZXInKS52YWwoKSkpIHtcbiAgICAgICQoXCIucGF5bWVudC1lcnJvcnNcIikudGV4dChcIkludmFsaWQgY2FyZCBudW1iZXJcIik7XG4gICAgICAkKFwiLnN1Ym1pdC1idXR0b25cIikucmVtb3ZlQXR0cihcImRpc2FibGVkXCIpO1xuICAgICAgcmV0dXJuIGZhbHNlO1xuICAgIH1cblxuICAgIGlmICghcGF5bWlsbC52YWxpZGF0ZUV4cGlyeShcbiAgICAgICQoJy5jYXJkLWV4cGlyeS1tb250aCcpLnZhbCgpLFxuICAgICAgJCgnLmNhcmQtZXhwaXJ5LXllYXInKS52YWwoKSlcbiAgICApIHtcbiAgICAgICQoXCIucGF5bWVudC1lcnJvcnNcIikudGV4dChcIkludmFsaWQgZXhwaXJhdGlvbiBkYXRlXCIpO1xuICAgICAgJChcIi5zdWJtaXQtYnV0dG9uXCIpLnJlbW92ZUF0dHIoXCJkaXNhYmxlZFwiKTtcbiAgICAgIHJldHVybiBmYWxzZTtcbiAgICB9XG5cbiAgICBjb25zb2xlLmxvZyhwYXltaWxsKVxuXG4gICAgcGF5bWlsbC5jcmVhdGVUb2tlbih7XG4gICAgICBudW1iZXI6ICAgICAgICAgJCgnLmNhcmQtbnVtYmVyJykudmFsKCksXG4gICAgICBleHBfbW9udGg6ICAgICAgJCgnLmNhcmQtZXhwaXJ5LW1vbnRoJykudmFsKCksXG4gICAgICBleHBfeWVhcjogICAgICAgJCgnLmNhcmQtZXhwaXJ5LXllYXInKS52YWwoKSxcbiAgICAgIGN2YzogICAgICAgICAgICAkKCcuY2FyZC1jdmMnKS52YWwoKSxcbiAgICAgIGNhcmRob2xkZXI6ICAgICAkKCcuY2FyZC1ob2xkZXJuYW1lJykudmFsKCksXG4gICAgICBhbW91bnRfaW50OiAgICAgJCgnLmNhcmQtYW1vdW50LWludCcpLnZhbCgpLCAvLyBJbnRlZ2VyIGUuZy4gXCI0OTAwXCIgZsODwrxyIDQ5LDAwIEVVUlxuICAgICAgY3VycmVuY3k6ICAgICAgICQoJy5jYXJkLWN1cnJlbmN5JykudmFsKCkgICAgLy8gSVNPIDQyMTcgei5CLiBcIkVVUlwiXG4gICAgICB9LCBQYXltaWxsUmVzcG9uc2VIYW5kbGVyKTtcblxuICAgIHJldHVybiBmYWxzZTtcbiAgfSk7XG59KTsiXX0=
