import { describe, it, expect, vi, beforeEach } from 'vitest';
import { mount } from '@vue/test-utils';
import ExamList from './ExamList.vue';
import api from '../services/api';

// Mock the API service
vi.mock('../services/api');

describe('ExamList.vue', () => {
    beforeEach(() => {
        vi.clearAllMocks();
    });

    it('renders "loading" state initially', async () => {
        api.getExams.mockReturnValue(new Promise(() => { })); // Never resolves
        const wrapper = mount(ExamList);
        expect(wrapper.text()).toContain('กำลังโหลด...');
    });

    it('renders empty state when no exams exist', async () => {
        api.getExams.mockResolvedValue({ data: [] });
        const wrapper = mount(ExamList);
        await vi.dynamicImportSettled(); // Wait for onMounted
        // Using a more robust check since we wait for the mock to resolve
        await new Promise(resolve => setTimeout(resolve, 0));
        expect(wrapper.text()).toContain('ยังไม่มีข้อสอบในระบบ');
    });

    it('renders a list of exams', async () => {
        const mockExams = [
            { id: 1, question: 'Q1', choices: [{ id: 1, choice_text: 'C1' }] }
        ];
        api.getExams.mockResolvedValue({ data: mockExams });
        const wrapper = mount(ExamList);
        await new Promise(resolve => setTimeout(resolve, 0));
        expect(wrapper.text()).toContain('Q1');
        expect(wrapper.text()).toContain('C1');
    });

    it('calls delete API when delete button is clicked', async () => {
        window.confirm = vi.fn(() => true);
        const mockExams = [{ id: 1, question: 'Q1', choices: [] }];
        api.getExams.mockResolvedValue({ data: mockExams });
        api.deleteExam.mockResolvedValue({});

        const wrapper = mount(ExamList);
        await new Promise(resolve => setTimeout(resolve, 0));

        await wrapper.find('.btn-danger').trigger('click');
        expect(api.deleteExam).toHaveBeenCalledWith(1);
    });
});
