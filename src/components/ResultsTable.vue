<template>
  <div class='parent flex-parent'>
    <div class='child flex-child'>
      <v2-table :data="list">
        <v2-table-column label="Nom" prop="name"></v2-table-column>
        <v2-table-column label="Temps" prop="laptime"></v2-table-column>
        <v2-table-column label="Voiture" prop="car"></v2-table-column>
      </v2-table>
    </div>
    <div class='child flex-child'>
      <v2-table :data="list">
        <v2-table-column label="Nom" prop="name"></v2-table-column>
        <v2-table-column label="Temps" prop="laptime"></v2-table-column>
        <v2-table-column label="Voiture" prop="car"></v2-table-column>
      </v2-table>
    </div>
    <div class='child flex-child'>
      <v2-table :data="list">
        <v2-table-column label="Nom" prop="name"></v2-table-column>
        <v2-table-column label="Temps" prop="laptime"></v2-table-column>
        <v2-table-column label="Voiture" prop="car"></v2-table-column>
      </v2-table>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'ResultsTable',
    data () {
      return {
        list: [{
          name: 'Tutu',
          laptime: '1\'28"50',
          car: 'Tutu mobile'
        }]
      }
    }
  }
</script>

<style>
.flex-parent {
  display: flex;
  border: 3px solid black;
}
.flex-child {
  flex: 1;
  border: 2px solid gold;
  border-radius: 5px;
}
</style>
