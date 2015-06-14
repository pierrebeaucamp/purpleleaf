$(document).ready(function() {
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
  
  $("#payment-form").on('submit', function(event) {
    // Deactivate submit button to avoid further clicks
    $('.submit-button').attr("disabled", "disabled");
    
    if (!paymill.validateCardNumber($('.card-number').val())) {
      $(".payment-errors").text("Invalid card number");
      $(".submit-button").removeAttr("disabled");
      return false;
    }

    if (!paymill.validateExpiry(
      $('.card-expiry-month').val(),
      $('.card-expiry-year').val())
    ) {
      $(".payment-errors").text("Invalid expiration date");
      $(".submit-button").removeAttr("disabled");
      return false;
    }

    paymill.createToken({
      number:         $('.card-number').val(),
      exp_month:      $('.card-expiry-month').val(),
      exp_year:       $('.card-expiry-year').val(),
      cvc:            $('.card-cvc').val(),
      cardholder:     $('.card-holdername').val(),
      amount_int:     $('.card-amount-int').val(), // Integer e.g. "4900" fÃ¼r 49,00 EUR
      currency:       $('.card-currency').val()    // ISO 4217 z.B. "EUR"
      }, PaymillResponseHandler);

    return false;
  });
});