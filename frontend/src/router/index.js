import { createRouter, createWebHistory } from 'vue-router';
import ExamList from '../views/ExamList.vue';
import AddExam from '../views/AddExam.vue';

const routes = [
    {
        path: '/',
        name: 'ExamList',
        component: ExamList,
    },
    {
        path: '/add',
        name: 'AddExam',
        component: AddExam,
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
