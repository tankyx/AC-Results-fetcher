<template>
  <div class="golddriver">
    <h1>Concours GOLD DRIVER</h1>
    <GoldDriverTable></GoldDriverTable>
  </div>
</template>

<script>
import GoldDriverTable from '@/components/GoldDriverTable.vue'

export default {
  name: 'GoldDriver',
  components: {
    GoldDriverTable
  }
}
</script>
