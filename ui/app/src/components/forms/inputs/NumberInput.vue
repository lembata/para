<script setup>
import { ref, defineModel, defineProps } from 'vue';
import InputNumber from 'primevue/inputnumber';
import { useI18n } from 'vue-i18n';
const { t } = useI18n(); 

const model = defineModel();
const props = defineProps(['text', 'placeholder', 'disabled',
  'ariaLabel', 'min', 'max', 'minFractionDigits']);

</script>

<template>
  <div class="form-input">
    <label class="form-label" for="grid-first-name">
      {{ $t(text) }}: {{ model }}
    </label>
    <InputNumber class="w-full" v-model="model" :disabled :aria-label :placeholder :min :max :minFractionDigits />
  </div>
</template>
