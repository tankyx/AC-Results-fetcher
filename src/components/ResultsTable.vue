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
      fetchResults() {

        console.log("Fetchin results");

       let url = 'http://127.0.0.1:8080/results';

        fetch(url)
        .then(
          function(response) {
            if (response.status !== 200) {
              console.log('Looks like there was a problem. Status Code: ' +
              response.status);
              return;
            }

            // Examine the text in the response
            response.json().then(function(data) {
              let parsed_data = JSON.parse(data);
              console.log(parsed_data.TrackName);
              // Peut pas modifier les result1 - 2 - 3 ici
            });
          }
        )
        .catch(err => { throw err });


        console.log("Ping")
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
