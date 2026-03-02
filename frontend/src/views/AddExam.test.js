import { describe, it, expect, vi, beforeEach } from 'vitest';
import { mount } from '@vue/test-utils';
import AddExam from './AddExam.vue';
import api from '../services/api';
import { createRouter, createWebHistory } from 'vue-router';

// Mock API
vi.mock('../services/api');

// Simple router mock
const router = createRouter({
    history: createWebHistory(),
    routes: [{ path: '/', component: { template: 'div' } }]
});

describe('AddExam.vue', () => {
    beforeEach(() => {
        vi.clearAllMocks();
    });

    it('shows alert if question is empty', async () => {
        const alertMock = vi.spyOn(window, 'alert').mockImplementation(() => { });
        const wrapper = mount(AddExam, { global: { plugins: [router] } });

        await wrapper.find('form').trigger('submit.prevent');
        expect(alertMock).toHaveBeenCalledWith('กรุณากรอกโจทย์');
    });

    it('calls createExam API with correct data', async () => {
        api.createExam.mockResolvedValue({});
        const wrapper = mount(AddExam, { global: { plugins: [router] } });

        await wrapper.find('textarea').setValue('Test Question?');
        const inputs = wrapper.findAll('input[type="text"]');
        await inputs[0].setValue('A1');
        await inputs[1].setValue('A2');
        await inputs[2].setValue('A3');
        await inputs[3].setValue('A4');

        await wrapper.find('form').trigger('submit.prevent');

        expect(api.createExam).toHaveBeenCalledWith({
            question: 'Test Question?',
            choices: ['A1', 'A2', 'A3', 'A4']
        });
    });
});
