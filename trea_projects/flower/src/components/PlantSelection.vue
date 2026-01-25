<script setup>
import { ref, watch } from 'vue'

// 接收父组件传递的属性
const props = defineProps({
  plants: {
    type: Array,
    required: true
  },
  selected: {
    type: Object,
    default: null
  }
})

// 定义事件
const emit = defineEmits(['update:selected'])

// 自定义植物名称
const customPlantName = ref('')

// 监听选中的植物变化
watch(() => props.selected, (newVal) => {
  if (newVal && newVal.isCustom) {
    customPlantName.value = newVal.name
  } else {
    customPlantName.value = ''
  }
})

// 处理植物图片上传成功
const handleImageUpload = (response, file, fileList) => {
  // 将图片转换为base64格式存储
  const reader = new FileReader()
  reader.onload = (e) => {
    // 创建一个新的植物对象，包含上传的图片
    const updatedPlant = {
      ...props.selected,
      image: e.target.result
    }
    emit('update:selected', updatedPlant)
  }
  reader.readAsDataURL(file.raw)
}

// 选择植物
const selectPlant = (plant) => {
  emit('update:selected', plant)
}

// 添加自定义植物
const addCustomPlant = () => {
  if (customPlantName.value.trim()) {
    const customPlant = {
      id: Date.now(),
      name: customPlantName.value.trim(),
      icon: '🌱',
      image: '',
      isCustom: true
    }
    emit('update:selected', customPlant)
  }
}
</script>

<template>
  <div class="plant-selection">
    <h3>选择植物</h3>
    
    <!-- 植物网格 -->
    <div class="plant-grid">
      <div 
        v-for="plant in plants" 
        :key="plant.id"
        class="plant-card"
        :class="{ selected: selected && selected.id === plant.id }"
        @click="selectPlant(plant)"
      >
        <div class="plant-icon">{{ plant.icon }}</div>
        <div class="plant-name">{{ plant.name }}</div>
      </div>
    </div>
    
    <!-- 自定义植物 -->
    <div class="custom-plant-section">
      <el-input 
        v-model="customPlantName" 
        placeholder="或输入自定义植物名称"
        clearable
        @keyup.enter="addCustomPlant"
      >
        <template #append>
          <el-button type="primary" @click="addCustomPlant">
            添加
          </el-button>
        </template>
      </el-input>
      
      <!-- 植物图片上传 -->
      <div class="image-upload" v-if="selected">
        <el-upload
          class="avatar-uploader"
          :show-file-list="false"
          :on-success="handleImageUpload"
          accept="image/*"
        >
          <img v-if="selected.image" :src="selected.image" class="plant-image" />
          <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
        </el-upload>
        <span class="upload-hint">点击上传植物图片</span>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* 植物选择区域样式 */
.plant-selection {
  margin-bottom: 25px;
}

.plant-selection h3 {
  color: #718096;
  margin-bottom: 15px;
  font-size: 1.1rem;
  font-weight: 500;
}

/* 植物网格 */
.plant-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(100px, 1fr));
  gap: 15px;
  margin-bottom: 20px;
}

/* 植物卡片 */
.plant-card {
  background: #f7fafc;
  border: 2px solid transparent;
  border-radius: 12px;
  padding: 15px;
  text-align: center;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
}

.plant-card:hover {
  border-color: #9f7aea;
  background: #edf2f7;
  transform: scale(1.05);
}

.plant-card.selected {
  border-color: #667eea;
  background: #e6fffa;
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.15);
}

.plant-icon {
  font-size: 2rem;
  margin-bottom: 5px;
}

.plant-name {
  font-size: 0.85rem;
  font-weight: 500;
  color: #4a5568;
}

/* 自定义植物区域 */
.custom-plant-section {
  display: flex;
  flex-direction: column;
  gap: 15px;
  padding: 15px;
  background: #f0fff4;
  border-radius: 10px;
  border: 1px dashed #9ae6b4;
}

/* 图片上传 */
.image-upload {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
  margin-top: 10px;
}

.plant-image {
  width: 100px;
  height: 100px;
  border-radius: 10px;
  object-fit: cover;
  border: 2px solid #e2e8f0;
  cursor: pointer;
  transition: all 0.3s ease;
}

.plant-image:hover {
  transform: scale(1.1);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.avatar-uploader-icon {
  font-size: 28px;
  color: #909399;
  background-color: #f5f7fa;
  width: 100px;
  height: 100px;
  line-height: 100px;
  text-align: center;
  border: 1px dashed #d9d9d9;
  border-radius: 10px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.avatar-uploader-icon:hover {
  border-color: #667eea;
  color: #667eea;
}

.upload-hint {
  font-size: 0.85rem;
  color: #6b7280;
}
</style>