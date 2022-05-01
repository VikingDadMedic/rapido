<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2020 David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
      <Card>
        {{ $t('settings_about__version') }} <span v-for="(st, key, idx) in status" :key="idx">{{st}}.</span><br><br>
        <div v-html="$t('settings_about__license')"></div><br>
        <div v-html="$t('settings_about__manual')"></div>
      </Card>
</template>

<script>
import { mapGetters } from 'vuex'
import axios from 'axios'
export default {
  name: 'admin-settings-about',

  mounted () {
    this.getStatus()
  },

  data () {
    return {
      status: {}
    }
  },

  computed: mapGetters({
    token: 'auth/token'
  }),

  methods: {
    // Get local system version
    async getStatus () {
      try {
        const { data } = await axios.get(process.env.API_URL + '/api/version', {headers: {Authorization: `Bearer ${this.token}`}})

        this.$set(this, 'status', data)
      } catch (e) {
        this.$awn.alert(this.$t('settings_about__version_file_reading_failed'))
      }
    }
  }
}
</script>
