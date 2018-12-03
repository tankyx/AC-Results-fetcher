<template>
  <div class='parent inline-flex-parent'>
    <div class='gdResultTable'>
      <v2-table :data="list" border>
        <v2-table-column label="Nom" prop="name"></v2-table-column>
        <v2-table-column label="Temps" prop="laptime"></v2-table-column>
        <v2-table-column label="Rang" prop="ranking" width="110">
          <template slot-scope="scope">
            <div class="medal">
              <img alt="Medal" src="../assets/Gold.png" height="35" width="35">
            </div>
          </template>
        </v2-table-column>
      </v2-table>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'GoldDriverTable',
    data () {
      return {
        list: [{
          name: 'Stig',
          laptime: '1\'28"50',
          ranking: '../assets/Gold.png',
        }]
      }
    }
  }
</script>

<style>
.inline-flex-parent {
  display: inline-flex;
  border: 3px solid black;
  border-radius: 5px;
}
.medal {
  display: block;
  margin-left: auto;
  margin-right: auto;
}
</style>
