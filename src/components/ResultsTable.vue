<template>
  <div class='parent flex-parent'>
    <div class='child flex-child'>
      <v2-table :lazy-load="true" :data="result1" border>
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

  function str_pad_left(string,pad,length) {
    return (new Array(length+1).join(pad)+string).slice(-length);
  }

  export default {
    name: 'ResultsTable',
    data () {
      return {
        result1: [],
        result2: [],
        result3: []
      }
    },
    async mounted() {
        console.log("Fetchin results");

       let url = '/results';
       try {
        const response = await axios.get(url);
        console.log(response.data);
        var results_data = JSON.parse(response.data);
        var results_list = []
        for (var i = 0; i < results_data.Result.length; i++)
        {
          var driver_res = results_data.Result[i];
          var minutes = Math.floor(driver_res.BestLap / 1000 / 60);
          var seconds = (driver_res.BestLap / 1000) - minutes * 60;

          console.log(minutes);
          console.log(seconds);

          var finalTime = str_pad_left(minutes,'0',2)+':'+new Array(2).join(seconds.toFixed(3));

          //Ca ne met pas à jour sur la page
          var tmp_res = { name: driver_res.DriverName,
                      laptime: finalTime,
                      car: driver_res.CarModel,
                      ranking: require('../assets/Gold.png')
                    };
          results_list.push(tmp_res);
          console.log(results_list); //Le log montre bien les changement, mais ca ne modifie pas sur la page
        }
       } catch (error) {
         console.log('rip: ', error)
          return;
       }
       this.result1 = results_list;
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
