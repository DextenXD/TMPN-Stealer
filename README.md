
<p align="center">
    <img src="./.github/assets/avatar.png" width=100  >
</p>



<h1 align="center">TMPN Stealer</h1>

<p align="center">Go-written Malware targeting Windows systems, extracting User Data from Discord, Browsers, Crypto Wallets and more, from every user on every disk. (PoC. For Educational Purposes only)</p>

---

## About the project

This proof of concept project demonstrates a "Discord-oriented" stealer implemented in Go. The malware operates on Windows systems and use fodhelper.exe technique for privileges elevation. By elevating privileges, the malware gains access to all user sessions on every disk

### Features:
- [x] - **antidebug**
  - [x] - **antivirus**
  - [x] - **antivm**
  - [x] - **browsers**
	  - Steals logins, cookies, credit cards, history, and download lists from 37 Chromium-based browsers.
 	  - Steals logins, cookies, history, and download lists from 10 Gecko browsers.
  - [x] - **clipper**
  - [x] - **commonfiles**
  - [x] - **discodes**
  - [x] - **discordinjection**
	  - Intercepts login, register, and 2FA login requests.
  	  - Captures backup codes requests.
  	  - Monitors email/password change requests.
 	  - Intercepts credit card/PayPal addition requests.
  	  - Blocks the use of QR codes for login.
 	  - Prevents requests to view devices.
  - [x] - **fakerror**
- [x] - **games**
- [x] - **hideconsole**
- [x] - **startup**
- [x] - **system**
- [x] - **tokens**
- [x] - **uacbypass**
- [x] - **wallets**
  - [x] - **walletsinjection**


## Disclaimer:

By installing and using this program, you agree to the following:

This program is provided for educational purposes only. Any actions taken based on the information provided are solely at your own risk. I, [Your Name], am not responsible for any consequences resulting from the use or misuse of this program. It is your responsibility to verify the accuracy and applicability of the information provided before taking any action. Always consult with appropriate professionals or experts if you have any doubts or concerns. By using this program, you agree to indemnify and hold harmless [Your Name] from any liability arising out of your use of the information herein.

