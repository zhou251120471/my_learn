<script setup>
import { ref, computed, onMounted } from 'vue'

// 接收父组件传递的属性
const props = defineProps({
  records: {
    type: Array,
    required: true
  }
})

// 当前显示的日期
const currentDate = ref(new Date())

// 选中的日期
const selectedDate = ref(new Date().toISOString().split('T')[0])

// 月份名称
const monthNames = [
  '一月', '二月', '三月', '四月', '五月', '六月',
  '七月', '八月', '九月', '十月', '十一月', '十二月'
]

// 星期名称
const weekNames = ['日', '一', '二', '三', '四', '五', '六']

// 计算当前月份的天数
const daysInMonth = computed(() => {
  const year = currentDate.value.getFullYear()
  const month = currentDate.value.getMonth() + 1
  return new Date(year, month, 0).getDate()
})

// 计算当前月份第一天是星期几
const firstDayOfMonth = computed(() => {
  const year = currentDate.value.getFullYear()
  const month = currentDate.value.getMonth()
  return new Date(year, month, 1).getDay()
})

// 计算当前月份的日历数据
const calendarDays = computed(() => {
  const days = []
  const year = currentDate.value.getFullYear()
  const month = currentDate.value.getMonth()
  
  // 添加上个月的天数
  const prevMonthDays = new Date(year, month, 0).getDate()
  for (let i = firstDayOfMonth.value - 1; i >= 0; i--) {
    days.push({
      day: prevMonthDays - i,
      date: new Date(year, month - 1, prevMonthDays - i).toISOString().split('T')[0],
      isCurrentMonth: false
    })
  }
  
  // 添加当前月的天数
  for (let i = 1; i <= daysInMonth.value; i++) {
    days.push({
      day: i,
      date: new Date(year, month, i).toISOString().split('T')[0],
      isCurrentMonth: true
    })
  }
  
  // 添加下个月的天数，使日历完整显示6行
  const remainingDays = 42 - days.length
  for (let i = 1; i <= remainingDays; i++) {
    days.push({
      day: i,
      date: new Date(year, month + 1, i).toISOString().split('T')[0],
      isCurrentMonth: false
    })
  }
  
  return days
})

// 获取当前月份的养护记录日期
const getCareDates = computed(() => {
  const year = currentDate.value.getFullYear()
  const month = currentDate.value.getMonth()
  
  return props.records
    .filter(record => {
      const recordDate = new Date(record.date)
      return recordDate.getFullYear() === year && recordDate.getMonth() === month
    })
    .map(record => record.date)
    .filter((date, index, self) => self.indexOf(date) === index) // 去重
})

// 检查某天是否有养护记录
const hasCareRecord = (date) => {
  return getCareDates.value.includes(date)
}

// 获取某天的养护记录数量
const getRecordCount = (date) => {
  return props.records.filter(record => record.date === date).length
}

// 切换到上个月
const prevMonth = () => {
  currentDate.value = new Date(currentDate.value.getFullYear(), currentDate.value.getMonth() - 1, 1)
}

// 切换到下个月
const nextMonth = () => {
  currentDate.value = new Date(currentDate.value.getFullYear(), currentDate.value.getMonth() + 1, 1)
}

// 获取当前月份标题
const currentMonthTitle = computed(() => {
  return `${currentDate.value.getFullYear()}年${monthNames[currentDate.value.getMonth()]}`
})

// 选择日期
const selectDate = (date) => {
  selectedDate.value = date
}

// 获取选中日期的养护记录
const selectedDateRecords = computed(() => {
  return props.records.filter(record => record.date === selectedDate.value)
})

// 格式化日期显示
const formatDateDisplay = (dateStr) => {
  const date = new Date(dateStr)
  return `${date.getFullYear()}年${date.getMonth() + 1}月${date.getDate()}日`
}

// 初始化时选中当天
onMounted(() => {
  selectedDate.value = new Date().toISOString().split('T')[0]
})
</script>

<template>
  <div class="calendar-container">
    <!-- 日历头部 -->
    <div class="calendar-header">
      <el-button 
        type="default" 
        @click="prevMonth" 
        class="month-nav-btn"
        icon="el-icon-arrow-left"
      ></el-button>
      <h3 class="month-title">{{ currentMonthTitle }}</h3>
      <el-button 
        type="default" 
        @click="nextMonth" 
        class="month-nav-btn"
        icon="el-icon-arrow-right"
      ></el-button>
    </div>
    
    <!-- 日历星期标题 -->
    <div class="calendar-weekdays">
      <div 
        v-for="(weekday, index) in weekNames" 
        :key="index"
        class="weekday"
      >
        {{ weekday }}
      </div>
    </div>
    
    <!-- 日历网格 -->
    <div class="calendar-grid">
      <div 
        v-for="(day, index) in calendarDays" 
        :key="index"
        class="calendar-day"
        :class="{
          'other-month': !day.isCurrentMonth,
          'has-record': hasCareRecord(day.date),
          'selected': selectedDate === day.date
        }"
        @click="selectDate(day.date)"
      >
        <span class="day-number">{{ day.day }}</span>
        <span 
          v-if="hasCareRecord(day.date)" 
          class="record-indicator"
        >
          {{ getRecordCount(day.date) }}
        </span>
      </div>
    </div>
    
    <!-- 当日养护记录详情 -->
    <div class="day-details">
      <h4 class="detail-title">{{ formatDateDisplay(selectedDate) }} 养护记录</h4>
      
      <div v-if="selectedDateRecords.length > 0" class="records-list">
        <el-collapse>
          <el-collapse-item 
            v-for="record in selectedDateRecords" 
            :key="record.id"
            :title="`${record.plantIcon} ${record.plantName} - ${record.action}`"
          >
            <div class="record-item">
              <div class="record-time">时间：{{ record.time }}</div>
              <div v-if="record.status" class="record-status">
                <strong>状态：</strong>{{ record.status }}
              </div>
              <div v-if="record.photo" class="record-photo">
                <strong>照片：</strong>
                <el-image 
                  :src="record.photo" 
                  :preview-src-list="[record.photo]" 
                  style="width: 100%; height: auto; margin-top: 10px; border-radius: 8px;"
                ></el-image>
              </div>
            </div>
          </el-collapse-item>
        </el-collapse>
      </div>
      
      <div v-else class="no-records">
        <el-empty description="当天没有养护记录"></el-empty>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* 日历容器 */
.calendar-container {
  width: 100%;
}

/* 日历头部 */
.calendar-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.month-nav-btn {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  padding: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
}

.month-nav-btn:hover {
  background-color: #f0f0f0;
  transform: scale(1.1);
}

.month-title {
  margin: 0;
  color: #5a67d8;
  font-size: 1.3rem;
  font-weight: 600;
}

/* 星期标题 */
.calendar-weekdays {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 5px;
  margin-bottom: 10px;
}

.weekday {
  text-align: center;
  padding: 10px 0;
  font-weight: 600;
  color: #718096;
  font-size: 0.9rem;
}

/* 日历网格 */
.calendar-grid {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 5px;
  margin-bottom: 25px;
}

/* 日历日期 */
.calendar-day {
  aspect-ratio: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.3s ease;
  font-weight: 500;
  position: relative;
  background: #f8fafc;
  border: 1px solid #e2e8f0;
}

.calendar-day:hover {
  background: #edf2f7;
  transform: scale(1.05);
}

.calendar-day.other-month {
  color: #d1d5db;
  background: #fafafa;
  cursor: default;
}

.calendar-day.other-month:hover {
  transform: none;
  background: #fafafa;
}

.calendar-day.has-record {
  background: #e6fffa;
  color: #2b6cb0;
  font-weight: 600;
}

.calendar-day.has-record:hover {
  background: #b2f5ea;
}

.calendar-day.selected {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
}

.calendar-day.selected:hover {
  transform: scale(1.05);
}

/* 日期数字 */
.day-number {
  font-size: 1rem;
}

/* 记录指示器 */
.record-indicator {
  position: absolute;
  bottom: 3px;
  background: #3182ce;
  color: white;
  font-size: 0.7rem;
  font-weight: 600;
  width: 16px;
  height: 16px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.calendar-day.selected .record-indicator {
  background: white;
  color: #667eea;
}

/* 当日详情 */
.day-details {
  margin-top: 25px;
  padding: 20px;
  background: #f0fff4;
  border-radius: 12px;
  border: 1px solid #c6f6d5;
}

.detail-title {
  margin: 0 0 15px 0;
  color: #2f855a;
  font-size: 1.1rem;
  font-weight: 600;
}

/* 记录列表 */
.records-list {
  max-height: 300px;
  overflow-y: auto;
}

/* 记录项 */
.record-item {
  padding: 15px;
  background: white;
  border-radius: 8px;
  margin-top: 10px;
  border-left: 3px solid #48bb78;
}

.record-time {
  color: #718096;
  font-size: 0.9rem;
  margin-bottom: 8px;
}

.record-status {
  color: #4a5568;
  line-height: 1.6;
  margin-bottom: 10px;
}

.record-photo {
  margin-top: 10px;
}

/* 无记录提示 */
.no-records {
  text-align: center;
  padding: 30px 0;
}

/* 滚动条样式 */
.records-list::-webkit-scrollbar {
  width: 6px;
}

.records-list::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

.records-list::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 3px;
}

.records-list::-webkit-scrollbar-thumb:hover {
  background: #a1a1a1;
}
</style>