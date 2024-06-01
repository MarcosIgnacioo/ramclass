import SubjectClass from "../classes/Subject"

export const filterSubjects = (subjects: Object[], subjectName: string, semester: number) => {
 // Fix pocho
 if (semester === 10) {
  return subjects.filter(subject => ((subject as SubjectClass).semester === 0))
 }
 subjects = subjects.filter(subject => (((subject as SubjectClass).subject_name.toLowerCase()).includes(subjectName.toLowerCase())))
 subjects = (semester === 0) ? subjects : subjects.filter(subject => (((subject as SubjectClass).semester) === semester))
 console.log(semester === 10)
 console.log(semester)
 return subjects
}

export const filter = () => {

}


