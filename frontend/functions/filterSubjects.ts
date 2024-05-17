import SubjectClass from "../classes/Subject"

export const filterSubjects = (subjects: Object[], subjectName: string, semester: number) => {
 subjects = subjects.filter(subject => (((subject as SubjectClass).subject_name.toLowerCase()).includes(subjectName.toLowerCase())))
 subjects = (semester === 0) ? subjects : subjects.filter(subject => (((subject as SubjectClass).semester) === semester))
 return subjects
}

export const filter = () => {

}


