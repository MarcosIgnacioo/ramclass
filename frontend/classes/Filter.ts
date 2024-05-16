
export default interface Filter {
 semester: number
 setSemester: React.Dispatch<React.SetStateAction<number>>
 subjectName: string
 setSubjectName: React.Dispatch<React.SetStateAction<string>>
}
