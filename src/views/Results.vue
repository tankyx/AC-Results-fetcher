<template>
  <div class="results">
    <h1>This is the results page</h1>
    <ResultsTable></ResultsTable>
  </div>
</template>

<script>
import ResultsTable from '@/components/ResultsTable.vue'

export default {
  name: 'Results',
  components: {
    ResultsTable
  }
}
</script>
