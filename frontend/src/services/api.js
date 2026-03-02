import axios from 'axios';

const apiClient = axios.create({
    baseURL: 'http://localhost:3000/api',
    headers: {
        'Content-Type': 'application/json',
    },
});

export default {
    getExams() {
        return apiClient.get('/exams');
    },
    createExam(exam) {
        return apiClient.post('/exams', exam);
    },
    deleteExam(id) {
        return apiClient.delete(`/exams/${id}`);
    },
};
