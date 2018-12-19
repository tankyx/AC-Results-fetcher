<template>
  <div class='parent flex-parent'>
    <div class='child flex-child'>
      <v2-table :data="result1" border>
        <v2-table-column label="Nom" prop="name"></v2-table-column>
        <v2-table-column label="Temps" prop="laptime"></v2-table-column>
        <v2-table-column label="Voiture" prop="car"></v2-table-column>
        <v2-table-column label="Rang" prop="ranking">
          <template slot-scope="scope">
            <div class="medal">
              <img alt="Medal" v-bind:src="scope.row.ranking" height="35" width="35">
            </div>
          </template>
        </v2-table-column>
      </v2-table>
    </div>
    <div class='child flex-child'>
      <v2-table :data="result2" border>
        <v2-table-column label="Nom" prop="name"></v2-table-column>
        <v2-table-column label="Temps" prop="laptime"></v2-table-column>
        <v2-table-column label="Voiture" prop="car"></v2-table-column>
        <v2-table-column label="Rang" prop="ranking">
          <template slot-scope="scope">
            <div class="medal">
              <img alt="Medal" v-bind:src="scope.row.ranking" height="35" width="35">
            </div>
          </template>
        </v2-table-column>
      </v2-table>
    </div>
    <div class='child flex-child'>
      <v2-table :data="result3" border>
        <v2-table-column label="Nom" prop="name"></v2-table-column>
        <v2-table-column label="Temps" prop="laptime"></v2-table-column>
        <v2-table-column label="Voiture" prop="car"></v2-table-column>
        <v2-table-column label="Rang" prop="ranking">
          <template slot-scope="scope">
            <div class="medal">
              <img alt="Medal" v-bind:src="scope.row.ranking" height="35" width="35">
            </div>
          </template>
        </v2-table-column>
      </v2-table>
    </div>
  </div>
</template>

<script>
  import axios from 'axios';
  export default {
    name: 'ResultsTable',
    data () {
      return {
        result1: [],
        result2: [],
        result3: []
      }
    },
    created: function()
    {
      this.fetchResults();
    },
    methods: {
      async fetchResults() {

        console.log("Fetchin results");

       let url = '/results';
       try {
        console.log("Wow");
        this.result1[0] = { name: 'Tutu',
                    laptime: '1\'28"50',
                    car: 'Tata Nano',
                    ranking: require('../assets/Gold.png')
                  };
        const response = await axios.get(url);
       } catch (error) {
         console.log('rip: ', error)
          return;
       }
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
