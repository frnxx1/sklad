// Переменные

var log = document.getElementById("Sign-in")
var reg = document.getElementById("Sign-up")

var logBtn = document.getElementById("log-btn")
var regBtn = document.getElementById("reg-btn")

var opnLog = document.getElementById("open-log")
var opnReg = document.getElementById("open-reg")



// Функция открытия формы регистрации

opnReg.addEventListener('click', ()=> {
	log.style.display = "none";
	reg.style.display = "block";
})



// Функция открытия формы входа

opnLog.addEventListener('click', ()=> {
	reg.style.display = "none";
	log.style.display = "block";
})







