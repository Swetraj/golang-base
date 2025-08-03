package emails

var PasswordResetTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Password Reset</title>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600&display=swap" rel="stylesheet">
    <script src="https://cdn.tailwindcss.com"></script>
    <style>
        body {
            font-family: 'Inter', sans-serif;
            background-color: #f7f7f7;
            margin: 0;
            padding: 0;
            -webkit-text-size-adjust: 100%;
            -ms-text-size-adjust: 100%;
        }
        .email-container {
            max-width: 600px;
            margin: 20px auto;
            background-color: #ffffff;
            border-radius: 12px;
            overflow: hidden;
            box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
        }
        .email-header {
            background-color: #4f46e5; /* Indigo 600 */
            color: #ffffff;
            padding: 32px;
            text-align: center;
            border-top-left-radius: 12px;
            border-top-right-radius: 12px;
        }
        .email-body {
            padding: 32px;
            color: #333333;
            line-height: 1.6;
        }
        .email-footer {
            background-color: #f0f0f0; /* Gray 200 */
            padding: 24px 32px;
            text-align: center;
            font-size: 14px;
            color: #666666;
            border-bottom-left-radius: 12px;
            border-bottom-right-radius: 12px;
        }
        .button {
            display: inline-block;
            background-color: #4f46e5; /* Indigo 600 */
            color: #ffffff;
            padding: 12px 24px;
            border-radius: 8px;
            text-decoration: none;
            font-weight: 600;
            margin-top: 24px;
            transition: background-color 0.3s ease;
        }
        .button:hover {
            background-color: #4338ca; /* Indigo 700 */
        }
        a {
            color: #4f46e5; /* Indigo 600 */
            text-decoration: none;
        }
        a:hover {
            text-decoration: underline;
        }
    </style>
</head>
<body class="bg-gray-100 p-4 sm:p-6 md:p-8">
    <div class="email-container shadow-lg">
        <!-- Email Header -->
        <div class="email-header">
            <h1 class="text-3xl font-bold">Password Reset Request</h1>
        </div>

        <!-- Email Body -->
        <div class="email-body">
            <p class="mb-4">Dear User,</p>
            <p class="mb-4">We received a request to reset the password for your account.</p>
            <p class="mb-4">To reset your password, please click on the link below:</p>
            <div class="text-center">
                <a href="{{RESET_LINK}}" class="button">Reset Your Password</a>
            </div>
            <p class="mt-8 text-sm text-gray-600">This link will expire in [e.g., 24 hours] for security reasons. If you did not request a password reset, please ignore this email. Your password will remain unchanged.</p>
            <p class="mt-4 text-sm text-gray-600">For your security, please do not share this link with anyone.</p>
            <p class="mt-6">If you have any questions, please contact our support team.</p>
            <p class="mt-8">Thank you,</p>
            <p class="mt-1">The [Your Company Name] Team</p>
        </div>

        <!-- Email Footer -->
        <div class="email-footer">
            <p>&copy; 2025 [Your Company Name]. All rights reserved.</p>
            <p>
                <a href="#" class="text-indigo-600 hover:underline">Privacy Policy</a> |
                <a href="#" class="text-indigo-600 hover:underline">Terms of Service</a>
            </p>
        </div>
    </div>
</body>
</html>

`
