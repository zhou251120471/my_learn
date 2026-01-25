<script setup>
import { ref, onMounted } from 'vue'
import PlantSelection from './components/PlantSelection.vue'
import CareForm from './components/CareForm.vue'
import Calendar from './components/Calendar.vue'
import Stats from './components/Stats.vue'

// 植物列表数据
const plants = ref([
  { id: 1, name: '绿萝', icon: '🌿', image: '' },
  { id: 2, name: '多肉', icon: '🍀', image: '' },
  { id: 3, name: '仙人掌', icon: '🌵', image: '' },
  { id: 4, name: '吊兰', icon: '🌱', image: '' },
  { id: 5, name: '君子兰', icon: '🌸', image: '' },
  { id: 6, name: '芦荟', icon: '🌿', image: '' }
])

// 选中的植物
const selectedPlant = ref(null)

// 养护记录数据
const careRecords = ref([])

// 预生成的示例数据
const sampleData = [
  {
    id: 1,
    plantId: 1,
    plantName: '绿萝',
    plantIcon: '🌿',
    action: '浇水',
    date: new Date(Date.now() - 86400000).toISOString().split('T')[0], // 昨天
    time: new Date(Date.now() - 86400000).toTimeString().split(' ')[0],
    status: '植物生长良好，叶片翠绿',
    photo: ''
  },
  {
    id: 2,
    plantId: 2,
    plantName: '多肉',
    plantIcon: '🍀',
    action: '浇水',
    date: new Date(Date.now() - 172800000).toISOString().split('T')[0], // 前天
    time: new Date(Date.now() - 172800000).toTimeString().split(' ')[0],
    status: '多肉叶片饱满，颜色鲜艳',
    photo: ''
  }
]

// 保存记录
const saveRecord = (record) => {
  // 生成唯一ID
  record.id = Date.now()
  // 添加到记录列表
  careRecords.value.unshift(record)
  // 保存到本地存储
  localStorage.setItem('careRecords', JSON.stringify(careRecords.value))
}

// 初始化数据
onMounted(() => {
  // 从本地存储读取数据
  const savedRecords = localStorage.getItem('careRecords')
  if (savedRecords) {
    careRecords.value = JSON.parse(savedRecords)
  } else {
    // 如果没有数据，使用预生成数据
    careRecords.value = sampleData
    localStorage.setItem('careRecords', JSON.stringify(sampleData))
  }
})
</script>

<template>
  <div class="app-container">
    <!-- 头部 -->
    <header class="app-header">
      <h1>🌱 花草养护记录</h1>
      <p>轻松记录每一次养护，让植物茁壮成长</p>
    </header>

    <!-- 主要内容 -->
    <main class="app-main">
      <!-- 左侧：养护记录 -->
      <section class="record-section">
        <el-card shadow="hover">
          <template #header>
            <h2>今日养护</h2>
          </template>
          
          <!-- 植物选择 -->
          <PlantSelection 
            :plants="plants" 
            v-model:selected="selectedPlant" 
          />
          
          <!-- 养护表单 -->
          <CareForm 
            :selected-plant="selectedPlant" 
            @save-record="saveRecord" 
          />
        </el-card>
      </section>

      <!-- 中间：养护日历 -->
      <section class="calendar-section">
        <el-card shadow="hover">
          <template #header>
            <h2>养护日历</h2>
          </template>
          
          <Calendar :records="careRecords" />
        </el-card>
      </section>

      <!-- 右侧：数据统计 -->
      <section class="stats-section">
        <el-card shadow="hover">
          <template #header>
            <h2>养护统计</h2>
          </template>
          
          <Stats :records="careRecords" :plants="plants" />
        </el-card>
      </section>
    </main>
  </div>
</template>

<style scoped>
/* 全局样式 */
.app-container {
  max-width: 1400px;
  margin: 0 auto;
  padding: 20px;
  background: linear-gradient(135deg, #f5f7fa 0%, #e8f0f2 100%);
  min-height: 100vh;
}

/* 头部样式 */
.app-header {
  text-align: center;
  margin-bottom: 30px;
  padding: 20px;
  background: rgba(255, 255, 255, 0.8);
  border-radius: 15px;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.05);
  backdrop-filter: blur(10px);
}

.app-header h1 {
  font-size: 2.5rem;
  color: #6b7280;
  margin-bottom: 10px;
  font-weight: 600;
}

.app-header p {
  color: #9ca3af;
  font-size: 1.1rem;
}

/* 主要内容区域 */
.app-main {
  display: grid;
  grid-template-columns: 1fr;
  gap: 30px;
}

/* 卡片样式 */
el-card {
  border-radius: 15px;
  overflow: hidden;
  transition: all 0.3s ease;
}

el-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.08);
}

el-card ::v-deep .el-card__header {
  background: #f8fafc;
  border-bottom: 1px solid #e2e8f0;
  padding: 20px;
}

el-card ::v-deep .el-card__body {
  padding: 20px;
}

el-card h2 {
  color: #5a67d8;
  margin: 0;
  font-size: 1.5rem;
  font-weight: 600;
}

/* 响应式设计 */
@media (min-width: 768px) {
  .app-main {
    grid-template-columns: 1fr 1fr;
  }
  
  .stats-section {
    grid-column: span 2;
  }
}

@media (min-width: 1024px) {
  .app-main {
    grid-template-columns: 1fr 1fr 1fr;
  }
  
  .stats-section {
    grid-column: span 1;
  }
}
</style>
