<script setup>
import { ref, computed, onMounted } from 'vue'

// 接收父组件传递的属性
const props = defineProps({
  records: {
    type: Array,
    required: true
  },
  plants: {
    type: Array,
    required: true
  }
})

// 总植物数
const totalPlants = computed(() => {
  // 从记录中提取所有不同的植物ID，包括自定义植物
  const plantIds = new Set()
  
  // 添加预设植物
  props.plants.forEach(plant => plantIds.add(plant.id))
  
  // 添加记录中的自定义植物
  props.records.forEach(record => {
    if (record.plantId && !plantIds.has(record.plantId)) {
      plantIds.add(record.plantId)
    }
  })
  
  return plantIds.size
})

// 本月浇水次数
const monthlyWaterCount = computed(() => {
  const now = new Date()
  const year = now.getFullYear()
  const month = now.getMonth()
  
  return props.records.filter(record => {
    const recordDate = new Date(record.date)
    return (
      recordDate.getFullYear() === year &&
      recordDate.getMonth() === month &&
      record.action === '浇水'
    )
  }).length
})

// 本月养护次数
const monthlyCareCount = computed(() => {
  const now = new Date()
  const year = now.getFullYear()
  const month = now.getMonth()
  
  return props.records.filter(record => {
    const recordDate = new Date(record.date)
    return (
      recordDate.getFullYear() === year &&
      recordDate.getMonth() === month
    )
  }).length
})

// 最近7天养护次数
const weeklyCareCount = computed(() => {
  const now = new Date()
  const sevenDaysAgo = new Date(now.getTime() - 7 * 24 * 60 * 60 * 1000)
  
  return props.records.filter(record => {
    const recordDate = new Date(record.date)
    return recordDate >= sevenDaysAgo && recordDate <= now
  }).length
})

// 最近的养护记录（最多5条）
const recentRecords = computed(() => {
  return [...props.records]
    .sort((a, b) => {
      // 先按日期排序，再按时间排序
      const dateCompare = new Date(b.date) - new Date(a.date)
      if (dateCompare !== 0) return dateCompare
      return b.time.localeCompare(a.time)
    })
    .slice(0, 5)
})

// 统计卡片数据
const statCards = [
  {
    title: '🌱 总植物数',
    value: totalPlants,
    color: '#48bb78',
    bgColor: 'linear-gradient(135deg, #48bb78 0%, #38a169 100%)'
  },
  {
    title: '💧 本月浇水',
    value: monthlyWaterCount,
    color: '#4299e1',
    bgColor: 'linear-gradient(135deg, #4299e1 0%, #3182ce 100%)'
  },
  {
    title: '📅 本月养护',
    value: monthlyCareCount,
    color: '#9f7aea',
    bgColor: 'linear-gradient(135deg, #9f7aea 0%, #805ad5 100%)'
  },
  {
    title: '🔥 最近7天',
    value: weeklyCareCount,
    color: '#f6ad55',
    bgColor: 'linear-gradient(135deg, #f6ad55 0%, #ed8936 100%)'
  }
]
</script>

<template>
  <div class="stats-container">
    <!-- 统计卡片 -->
    <div class="stats-grid">
      <div 
        v-for="(card, index) in statCards" 
        :key="index"
        class="stat-card"
        :style="{ background: card.bgColor }"
      >
        <h3 class="card-title">{{ card.title }}</h3>
        <p class="card-value">{{ card.value }}</p>
      </div>
    </div>
    
    <!-- 最近养护记录 -->
    <div class="recent-records-section">
      <h4 class="section-title">最近养护记录</h4>
      
      <div v-if="recentRecords.length > 0" class="recent-records-list">
        <el-timeline>
          <el-timeline-item 
            v-for="record in recentRecords" 
            :key="record.id"
            :timestamp="`${record.date} ${record.time}`"
            placement="top"
          >
            <div class="timeline-content">
              <div class="plant-info">
                <span class="plant-icon">{{ record.plantIcon }}</span>
                <span class="plant-name">{{ record.plantName }}</span>
              </div>
              <div class="action-info">
                <el-tag :type="record.action === '浇水' ? 'primary' : record.action === '施肥' ? 'success' : 'warning'">
                  {{ record.action }}
                </el-tag>
              </div>
              <div v-if="record.status" class="status-info">
                {{ record.status }}
              </div>
              <div v-if="record.photo" class="photo-info">
                <el-image 
                  :src="record.photo" 
                  :preview-src-list="[record.photo]"
                  style="width: 60px; height: 60px; border-radius: 8px; cursor: pointer; margin-top: 5px;"
                ></el-image>
              </div>
            </div>
          </el-timeline-item>
        </el-timeline>
      </div>
      
      <div v-else class="no-recent-records">
        <el-empty description="暂无养护记录"></el-empty>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* 统计容器 */
.stats-container {
  width: 100%;
}

/* 统计卡片网格 */
.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: 20px;
  margin-bottom: 30px;
}

/* 统计卡片 */
.stat-card {
  padding: 25px;
  border-radius: 16px;
  text-align: center;
  transition: all 0.3s ease;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
  color: white;
}

.stat-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.15);
}

.card-title {
  margin: 0 0 10px 0;
  font-size: 0.95rem;
  font-weight: 500;
  opacity: 0.9;
}

.card-value {
  margin: 0;
  font-size: 2.5rem;
  font-weight: 700;
  line-height: 1;
}

/* 最近记录区域 */
.recent-records-section {
  margin-top: 30px;
  padding: 20px;
  background: #f7fafc;
  border-radius: 12px;
  border: 1px solid #e2e8f0;
}

.section-title {
  margin: 0 0 20px 0;
  color: #4a5568;
  font-size: 1.1rem;
  font-weight: 600;
}

/* 最近记录列表 */
.recent-records-list {
  max-height: 300px;
  overflow-y: auto;
}

/* 时间线内容 */
.timeline-content {
  background: white;
  padding: 15px;
  border-radius: 10px;
  border-left: 3px solid #667eea;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
}

.plant-info {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 8px;
}

.plant-icon {
  font-size: 1.2rem;
}

.plant-name {
  font-weight: 600;
  color: #4a5568;
}

.action-info {
  margin-bottom: 8px;
}

.status-info {
  color: #718096;
  font-size: 0.9rem;
  line-height: 1.5;
  margin-bottom: 8px;
}

/* 无记录提示 */
.no-recent-records {
  text-align: center;
  padding: 30px 0;
}

/* 滚动条样式 */
.recent-records-list::-webkit-scrollbar {
  width: 6px;
}

.recent-records-list::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

.recent-records-list::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 3px;
}

.recent-records-list::-webkit-scrollbar-thumb:hover {
  background: #a1a1a1;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .stats-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 15px;
  }
  
  .stat-card {
    padding: 20px;
  }
  
  .card-value {
    font-size: 2rem;
  }
}

@media (max-width: 480px) {
  .stats-grid {
    grid-template-columns: 1fr;
  }
}
</style>