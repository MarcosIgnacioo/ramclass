export const BASE_PATH = window.location.origin + "/"

export const SemestersMap = {
 0: "optative",
 1: "first",
 2: "second",
 3: "third",
 4: "fourth",
 5: "fifth",
 6: "sixth",
 7: "seventh",
 8: "eighth",
 9: "ninth",
 10: "tenth",
}

export const semesters = ["Todos", "Primero", "Segundo", "Tercero", "Cuarto", "Quinto", "Sexto", "Séptimo", "Octavo", "Noveno", "Optativas"]

export function deactivateNavBar() {
 const navBarLinks = document.querySelectorAll(".nav-links ")
 navBarLinks.forEach(link => {
  link.classList.add("disabled-link")
 })
}

export function reactivateNavbar() {
 const navBarLinks = document.querySelectorAll(".nav-links ")
 navBarLinks.forEach(link => {
  link.classList.remove("disabled-link")
 })
}
