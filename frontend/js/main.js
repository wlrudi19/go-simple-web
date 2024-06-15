$(document).ready(function() {
    $('#loginForm').on('submit', function(e) {
        e.preventDefault();

        const email = $('#email').val();
        const password = $('#password').val();

        $.ajax({
            url: 'http://localhost:3012/api/users/login',
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify({ email: email, password: password }),
            success: function(response) {
                // Handle successful login here
                console.log('Login successful', response);
            },
            error: function(error) {
                // Handle login error here
                console.error('Login failed', error);
            }
        });
    });
});
