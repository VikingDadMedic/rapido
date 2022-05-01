<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <Main>
    <section class="section template-main">
      <div class="container-fluid h-100">
        <div class="row h-100 align-items-center">
          <div class="col-12 m-auto">
            <div class="text-center">
              <h2 class="mb-4">
                {{ 'Page Not Found' }}
              </h2>
              <div class="links">
                <router-link :to="{ name: 'Index' }">
                  {{ $t('settings_system__homepage') }}
                </router-link>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </Main>
</template>

<script>
export default {
  name: 'not-found',

  metaInfo () {
    return { title: 'Page Not Found' }
  },

  layout: 'default'
}
</script>

<style lang="stylus" scoped>
.template-main
  height: 100vh
</style>
