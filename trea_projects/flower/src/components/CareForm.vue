<script setup>
import { ref } from 'vue'
// 引入 Element Plus 消息提示组件
import { ElMessage } from 'element-plus'

// 接收父组件传递的属性
const props = defineProps({
  selectedPlant: {
    type: Object,
    default: null
  }
})

// 定义事件
const emit = defineEmits(['save-record'])

// 养护操作类型
const careActions = [
  { label: '💧 浇水', value: '浇水' },
  { label: '🌿 施肥', value: '施肥' },
  { label: '✂️ 修剪', value: '修剪' }
]

// 选中的养护操作
const selectedAction = ref('浇水')

// 状态描述
const statusDesc = ref('')

// 养护照片
const carePhoto = ref('')

// 处理养护操作选择
const handleActionSelect = (action) => {
  selectedAction.value = action
}

// 处理图片上传成功
const handlePhotoUpload = (response, file, fileList) => {
  // 将图片转换为base64格式存储
  const reader = new FileReader()
  reader.onload = (e) => {
    carePhoto.value = e.target.result
  }
  reader.readAsDataURL(file.raw)
}

// 保存养护记录
const saveCareRecord = () => {
  // 验证是否选择了植物
  if (!props.selectedPlant) {
    ElMessage.warning('请先选择或添加植物')
    return
  }

  // 获取当前日期和时间
  const now = new Date()
  const date = now.toISOString().split('T')[0]
  const time = now.toTimeString().split(' ')[0]

  // 构建养护记录
  const record = {
    plantId: props.selectedPlant.id,
    plantName: props.selectedPlant.name,
    plantIcon: props.selectedPlant.icon,
    action: selectedAction.value,
    date: date,
    time: time,
    status: statusDesc.value.trim(),
    photo: carePhoto.value
  }

  // 发送事件给父组件
  emit('save-record', record)

  // 重置表单
  resetForm()

  // 显示成功消息
  ElMessage.success('养护记录保存成功！')
}

// 重置表单
const resetForm = () => {
  selectedAction.value = '浇水'
  statusDesc.value = ''
  carePhoto.value = ''
}
</script>

<template>
  <div class="care-form">
    <h3>养护操作</h3>
    
    <!-- 操作按钮 -->
    <div class="action-buttons">
      <el-button
        v-for="action in careActions"
        :key="action.value"
        :type="selectedAction === action.value ? 'primary' : 'default'"
        @click="handleActionSelect(action.value)"
        class="action-btn"
      >
        {{ action.label }}
      </el-button>
    </div>
    
    <!-- 状态描述 -->
    <div class="status-section">
      <el-input
        v-model="statusDesc"
        type="textarea"
        :rows="3"
        placeholder="添加植物状态描述..."
        maxlength="200"
        show-word-limit
      ></el-input>
    </div>
    
    <!-- 照片上传 -->
    <div class="photo-upload-section">
      <h4>添加养护照片（可选）</h4>
      <el-upload
        class="care-photo-uploader"
        :show-file-list="false"
        :on-success="handlePhotoUpload"
        accept="image/*"
      >
        <img v-if="carePhoto" :src="carePhoto" class="care-photo" />
        <el-icon v-else class="photo-uploader-icon"><Plus /></el-icon>
      </el-upload>
      <p class="upload-tip">点击或拖拽图片到此处上传</p>
    </div>
    
    <!-- 保存按钮 -->
    <div class="submit-section">
      <el-button 
        type="success" 
        size="large" 
        @click="saveCareRecord"
        class="submit-btn"
        :disabled="!selectedPlant"
      >
        保存记录
      </el-button>
    </div>
  </div>
</template>

<style scoped>
/* 养护表单样式 */
.care-form {
  margin-top: 20px;
}

.care-form h3 {
  color: #718096;
  margin-bottom: 15px;
  font-size: 1.1rem;
  font-weight: 500;
}

.care-form h4 {
  color: #4a5568;
  margin-bottom: 10px;
  font-size: 0.95rem;
  font-weight: 500;
}

/* 操作按钮 */
.action-buttons {
  display: flex;
  gap: 12px;
  margin-bottom: 20px;
  flex-wrap: wrap;
}

.action-btn {
  flex: 1;
  min-width: 100px;
  border-radius: 25px;
  font-weight: 500;
  transition: all 0.3s ease;
}

.action-btn:hover {
  transform: translateY(-2px);
}

/* 状态描述区域 */
.status-section {
  margin-bottom: 20px;
}

.status-section ::v-deep .el-textarea__inner {
  border-radius: 10px;
  resize: vertical;
  transition: all 0.3s ease;
}

.status-section ::v-deep .el-textarea__inner:focus {
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

/* 照片上传区域 */
.photo-upload-section {
  margin-bottom: 25px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 15px;
}

.care-photo-uploader {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 150px;
  height: 150px;
  border-radius: 15px;
  overflow: hidden;
  cursor: pointer;
}

.care-photo {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border: 2px solid #e2e8f0;
  transition: all 0.3s ease;
}

.care-photo:hover {
  transform: scale(1.05);
}

.photo-uploader-icon {
  font-size: 40px;
  color: #909399;
  background-color: #f5f7fa;
  width: 100%;
  height: 100%;
  line-height: 150px;
  text-align: center;
  border: 2px dashed #d9d9d9;
  border-radius: 15px;
  transition: all 0.3s ease;
}

.photo-uploader-icon:hover {
  border-color: #667eea;
  color: #667eea;
  background-color: #f0f4ff;
}

.upload-tip {
  color: #909399;
  font-size: 0.85rem;
  margin: 0;
}

/* 提交按钮区域 */
.submit-section {
  display: flex;
  justify-content: center;
  margin-top: 30px;
}

.submit-btn {
  min-width: 150px;
  border-radius: 30px;
  padding: 12px 30px;
  font-size: 1rem;
  font-weight: 600;
  transition: all 0.3s ease;
}

.submit-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(64, 158, 255, 0.3);
}
</style>