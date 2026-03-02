<template>
  <div class="container">
    <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 2rem;">
      <h1>คลังข้อสอบ (Exam List)</h1>
      <button class="btn btn-primary" @click="$router.push('/add')">
        เพิ่มข้อสอบ (Add Exam)
      </button>
    </div>

    <div v-if="loading" class="card" style="text-align: center;"> กำลังโหลด... </div>
    
    <div v-else-if="exams.length === 0" class="card" style="text-align: center; color: var(--text-muted);">
      ยังไม่มีข้อสอบในระบบ
    </div>

    <div v-else>
      <div v-for="(exam, index) in exams" :key="exam.id" class="card">
        <div style="display: flex; justify-content: space-between; align-items: flex-start;">
          <div style="flex: 1;">
            <h3 style="margin-top: 0; margin-bottom: 1rem;">
              {{ index + 1 }}. {{ exam.question }}
            </h3>
            <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 0.75rem;">
              <div v-for="(choice, cIdx) in exam.choices" :key="choice.id" 
                   style="display: flex; align-items: center; gap: 0.5rem; padding: 0.5rem; border: 1px solid var(--border); border-radius: 6px;">
                <input type="radio" :name="'exam-' + exam.id" disabled>
                <span style="font-size: 0.9rem;">{{ choice.choice_text }}</span>
              </div>
            </div>
          </div>
          <button class="btn btn-danger" style="margin-left: 1rem;" @click="deleteExam(exam.id)">
            ลบ (Delete)
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import api from '../services/api';

const exams = ref([]);
const loading = ref(true);

const fetchExams = async () => {
  try {
    const response = await api.getExams();
    exams.value = response.data;
  } catch (error) {
    console.error('Failed to fetch exams:', error);
  } finally {
    loading.value = false;
  }
};

const deleteExam = async (id) => {
  if (confirm('คุณแน่ใจหรือไม่ว่าต้องการลบข้อสอบนี้?')) {
    try {
      await api.deleteExam(id);
      await fetchExams();
    } catch (error) {
      console.error('Failed to delete exam:', error);
    }
  }
};

onMounted(fetchExams);
</script>
