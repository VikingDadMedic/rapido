<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <div class="row">
    <div class="col-md-12">
      <Card>
        <div class="tab-content">
          <div class="tab-pane fade show active p-3" role="tabpanel" id="accordion-contactforms" aria-labelledby="accordion-contactforms-tab">
            <h5>
              {{ $t('settings_features__theme') }}
            </h5>

            <!-- Themes radio buttons -->
            <div class="form-group row">
              <span class="col-md-4 d-flex justify-content-end align-items-bottom">
                <input :checked="theme[0]" v-on:click="changeTheme(0)" name="theme" type="radio">
              </span>
              <label class="col-md-8 col-form-label">
                {{ $t('settings_features__default_theme') }}
              </label>
            </div>
            <div class="form-group row">
              <span class="col-md-4 d-flex justify-content-end align-items-bottom">
                <input :checked="theme[1]" v-on:click="changeTheme(1)" name="theme" type="radio">
              </span>
              <label class="col-md-8 col-form-label">
                {{ $t('settings_features__dark_theme') }}
              </label>
            </div>
            <div class="form-group row">
              <span class="col-md-4 d-flex justify-content-end align-items-bottom">
                <input :checked="theme[2]" v-on:click="changeTheme(2)" name="theme" type="radio">
              </span>
              <label class="col-md-8 col-form-label">
                {{ $t('settings_features__colourblind_theme') }}
              </label>
            </div>

            <!-- Save button -->
            <div class="form-group row">
              <span class="col-md-4 offset-md-4">
                <v-button :loading="theme.busy" @click="saveThemeSettings" class="btn btn-success">
                  {{ $t('button__save') }}
                </v-button>
              </span>
            </div>

          </div>
        </div>
      </Card>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import axios from 'axios'
import Form from 'vform'

export default {

  mounted () {
    const { data } = axios.get(process.env.API_URL + '/api/theme', {headers: {Authorization: `Bearer ${this.token}`}})
    this.theme[data[0].theme_id] = true
  },

  data () {
    return {
      processed: false,
      theme: new Form({
        0: true,
        1: false,
        2: false
      })
    }
  },

  computed: {
    ...mapGetters({
      token: 'auth/token'
    })
  },

  components: {
  },

  methods: {
    changeTheme (i) {
      Object.keys(this.theme).forEach(key => {
        if (!isNaN(key)) {
          this.theme[key] = false
        }
      })
      this.theme[i] = true
    },

    async saveThemeSettings () {
      let newThemeID = ''
      Object.keys(this.theme).forEach(key => {
        if (!isNaN(key) && this.theme[key]) {
          newThemeID = key
        }
      })
      try {
        const { data } = await axios.put(process.env.API_URL + '/api/theme?new=' + newThemeID, {}, {headers: {Authorization: `Bearer ${this.token}`}})
        this.$awn.success(this.$t(data.message))
      } catch (e) {
        this.$awn.alert(this.$t(e.response.data.message))
      }
    }

  }
}
</script>
