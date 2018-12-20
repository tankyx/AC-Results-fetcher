<template>
  <div class="results">
    <meta http-equiv="refresh" content="30" >
    <h1>RESULTATS</h1>
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
