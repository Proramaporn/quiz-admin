<template>
  <div class="container">
    <h1>เพิ่มข้อสอบ (Add Exam)</h1>
    
    <div class="card">
      <form @submit.prevent="saveExam">
        <div class="form-group">
          <label>โจทย์ (Question)</label>
          <textarea v-model="form.question" placeholder="พิมพ์โจทย์ที่นี่..." required></textarea>
        </div>

        <div style="margin-top: 1.5rem;">
          <label>ตัวเลือก (Choices)</label>
          <div v-for="i in 4" :key="i" class="choice-input">
            <span class="choice-num">{{ i }}</span>
            <input type="text" v-model="form.choices[i-1]" :placeholder="'ตัวเลือกที่ ' + i" required>
          </div>
        </div>

        <div style="display: flex; gap: 1rem; margin-top: 2rem; justify-content: flex-end;">
          <button type="button" class="btn btn-secondary" @click="$router.push('/')">ยกเลิก (Cancel)</button>
          <button type="submit" class="btn btn-primary" :disabled="submitting">
            {{ submitting ? 'กำลังบันทึก...' : 'บันทึก (Save)' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import api from '../services/api';

const router = useRouter();
const submitting = ref(false);

const form = reactive({
  question: '',
  choices: ['', '', '', '']
});

const saveExam = async () => {
  // Validation
  if (!form.question.trim()) {
    alert('กรุณากรอกโจทย์');
    return;
  }
  if (form.choices.some(c => !c.trim())) {
    alert('กรุณากรอกตัวเลือกให้ครบทั้ง 4 ข้อ');
    return;
  }

  submitting.value = true;
  try {
    const payload = {
      question: form.question,
      choices: form.choices // Backend expects [string, string, string, string]
    };
    await api.createExam(payload);
    router.push('/');
  } catch (error) {
    console.error('Failed to save exam:', error);
    alert('เกิดข้อผิดพลาดในการบันทึกข้อสอบ');
  } finally {
    submitting.value = false;
  }
};
</script>

<style scoped>
.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 600;
}

textarea {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid var(--border);
  border-radius: 8px;
  font-family: inherit;
  font-size: 1rem;
  min-height: 100px;
  box-sizing: border-box;
}

.choice-input {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-top: 0.75rem;
}

.choice-num {
  background: #e2e8f0;
  width: 28px;
  height: 28px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  font-size: 0.8rem;
  font-weight: bold;
}

input[type="text"] {
  flex: 1;
  padding: 0.75rem;
  border: 1px solid var(--border);
  border-radius: 8px;
  font-size: 1rem;
}

input:focus, textarea:focus {
  outline: none;
  border-color: var(--primary);
  ring: 2px var(--primary-light);
}
</style>
