<template>
  <div class='parent flex-parent'>
    <div class='child flex-child'>
      <v2-table :lazy-load="true" :data="result1" border>
        <v2-table-column label="Nom" prop="name"></v2-table-column>
        <v2-table-column label="Temps" prop="laptime"></v2-table-column>
        <v2-table-column label="Voiture" prop="car"></v2-table-column>
        <v2-table-column label="Rang" prop="ranking"></v2-table-column>
      </v2-table>
    </div>
    <div class='child flex-child'>
      <v2-table :data="result2" border>
        <v2-table-column label="Nom" prop="name"></v2-table-column>
        <v2-table-column label="Temps" prop="laptime"></v2-table-column>
        <v2-table-column label="Voiture" prop="car"></v2-table-column>
        <v2-table-column label="Rang" prop="ranking"></v2-table-column>
      </v2-table>
    </div>
    <div class='child flex-child'>
      <v2-table :data="result3" border>
        <v2-table-column label="Nom" prop="name"></v2-table-column>
        <v2-table-column label="Temps" prop="laptime"></v2-table-column>
        <v2-table-column label="Voiture" prop="car"></v2-table-column>
        <v2-table-column label="Rang" prop="ranking"></v2-table-column>
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

       let urls = ['/results', '/results2', '/results3'];
       try {
        for (var j = 0; j < 3; j++)
        {
          const response = await axios.get(urls[j]);
          console.log(response.data);
          var results_data = JSON.parse(response.data);
          var results_list = []
          for (var i = 0; i < results_data.Result.length; i++)
          {
            var driver_res = results_data.Result[i];
            var minutes = Math.floor(driver_res.BestLap / 1000 / 60);
            var seconds = (driver_res.BestLap / 1000) - minutes * 60;
            var finalTime = '';

            console.log(minutes);
            console.log(seconds);
            if (seconds < 10) {
              finalTime = str_pad_left(minutes,'0',2)+':'+str_pad_left(seconds.toFixed(3),'0',5);
            }
            else {
              finalTime = str_pad_left(minutes,'0',2)+':'+new Array(2).join(seconds.toFixed(3));
            }

            var tmp_res = { name: driver_res.DriverName,
                        laptime: finalTime,
                        car: driver_res.CarModel,
                        ranking: i + 1,
                      };
            results_list.push(tmp_res);
            console.log(results_list); //Le log montre bien les changement, mais ca ne modifie pas sur la page
          }
          if (j === 0) {
            this.result1 = results_list;
          }
          else if (j === 1) {
            this.result2 = results_list;
          }
          else {
            this.result3 = results_list;
          }
        }
        } catch (error) {
          console.log('rip: ', error)
          return;
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
