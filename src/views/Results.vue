<template>
  <div class="results">
    <h1>This is the results page</h1>
  </div>
  <div id="demo">
    <demo-grid
      :data="gridData"
      :columns="gridColumns"
      :filter-key="searchQuery">
    </demo-grid>
  </div>
</template>

<script type="text/x-template" id="grid-template">
  <table>
    <thead>
      <tr>
        <th v-for="key in columns"
          v-on:click="sortBy(key)"
          v-bind:class="{ active: sortKey == key }">
          {{ key , capitalize }}
          <span class="arrow" v-bind:class="sortOrders[key] > 0 ? 'asc' : 'dsc'">
          </span>
        </th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="entry in filteredData">
        <td v-for="key in columns">
          {{entry[key]}}
        </td>
      </tr>
    </tbody>
  </table>
</script>
