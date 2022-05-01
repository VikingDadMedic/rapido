<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <div :id="id" class="communication-contact-form">
    <div class="row">
      <div class="col-md-12">
       <form name="form" class="form-horizontal" role="form" @submit.prevent="sendMail" novalidate>

          <!-- Email adresses -->
          <div class="form-group row">
            <span class="col-md-12">
              <input id="contact-form-to" name="to" type="text" v-model="form.to" hidden>
            </span>
          </div>
          <div class="form-group row">
            <label for="contact-form-email" class="col-md-4 col-form-label text-md-right">
              {{ $t('your_email_address') }}
            </label>
            <span class="col-md-8">
              <input id="contact-form-email" name="email" class="form-control" type="text" v-model="form.email">
            </span>
          </div>

          <!-- Subject -->
          <div class="form-group row">
            <label for="contact-form-subject" class="col-md-4 col-form-label text-md-right">
              {{ $t('contact_form__subject') }}
            </label>
            <span class="col-md-8">
              <input id="contact-form-subject" name="subject" class="form-control" type="text" v-model="form.subject" >
            </span>
          </div>

          <!-- Message -->
          <div class="form-group row">
            <label for="contact-form-message" class="col-md-4 col-form-label text-md-right">
              {{ $t('contact_form__message') }}
            </label>
            <span class="col-md-8">
              <textarea id="contact-from-message" name="message" class="form-control" type="text" v-model="form.message"></textarea>
            </span>
          </div>

          <!-- Attachment(s) -->
          <div class="form-group row">
            <label for="contact-form-message" class="col-md-4 col-form-label text-md-right">
              {{ $t('contact_form__attachment') }}
            </label>
            <span class="col-md-8">
              <input type="file" ref="fileinput" name="files" class="form-control" @change="onFileChanged" multiple>
            </span>
          </div>
          <div class="form-group row">
            <label for="contact-form-human" class="col-md-4 col-form-label text-md-right">
              {{ $t('contact_form__captcha_title') }}
            </label>
            <span class="col-md-8">
              <input id="contact-form-human" name="human" class="form-control" type="text" v-model="form.human" >
            </span>
          </div>
          <div class="form-group row">
            <div class="col-md-4 offset-md-4">
              <button type="submit" class="btn btn-success" :class="{'btn-loading': processed}" :disabled="processed">
                <span class="fa fa-envelope" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('button__send_message') }}
              </button>
            </div>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
/**
 * Component documentation
 */
import { mapGetters } from 'vuex'
import Form from 'vform'
import axios from 'axios'
export default {

  name: 'ContactForm',

  mounted () {
    this.setSettings()
  },

  props: {
    id: {
      type: String,
      default: 'communication-contact-form'
    },
    to: {
      type: String,
      default: ''
    }
  },

  data: () => ({
    regexpPattern: /^.+@.+$/,
    processed: false,
    form: new Form({
      subject: null,
      to: null,
      email: null,
      message: null,
      files: [],
      human: 'test'
    })
  }),

  computed: mapGetters({
    contactforms: 'page/contactforms',
    access_level: 'page/access_level',
    token: 'auth/token',
    manifest: 'page/manifest'
  }),

  methods: {

    // onFileChanged() appends the file to form data
    onFileChanged (e) {
      for (let i in e.target.files) {
        if (e.target.files[i].type && e.target.files[i].len < 0) {
          let file = e.target.files[i]
          this.form.append('file[' + i + ']', file)
        }
      }
    },

    async sendMail () {
      this.appendData()
      if (!this.form.email || !this.regexpPattern.test(this.form.email.trim())) {
        this.$awn.alert(this.$t('your_email_address_is_invalid'))
      } else if (!this.form.to || !this.regexpPattern.test(this.form.to.trim())) {
        this.$awn.alert(this.$t('recipient_email_address_is_invalid'))
      } else if (!this.form.subject || this.form.subject.trim() === '') {
        this.$awn.alert(this.$t('contact_form__subject_field_required'))
      } else if (!this.form.message || this.form.message.trim() === '') {
        this.$awn.alert(this.$t('contact_form__message_field_required'))
      } else if (this.form.human) {
        this.$awn.alert(this.$t('contact_form__captcha_value_is_invalid'))
      } else {
        this.$validator.validateAll()
          .then(async (res) => {
            this.form.email = this.form.email.trim()
            this.form.to = this.form.to.trim()
            this.form.subject = this.form.subject.trim()
            this.form.message = this.form.message.trim()
            res = true
            if (res) {
              this.processed = true
              try {
                await axios.post(process.env.API_URL + '/api/email', this.form, {
                  headers: {
                    'Content-Type': 'multipart/form-data'
                  }
                })
                this.$awn.success(this.$t('contact_form__message_sent'))
                this.processed = false
              } catch (e) {
                this.$awn.alert(e.response.data.message)
                this.processed = false
              }
            } else {
              console.log(res)
            }
          })
      }
    },

    // Initialize FormData and populate to and captcha fields
    setSettings () {
      this.form = new FormData()
      this.form['to'] = this.to
      this.form['human'] = this.$t('contact_form__captcha_field')
    },

    // Append all data to the multipart Form
    appendData () {
      this.form.append('subject', this.form.subject)
      this.form.append('to', this.form.to)
      this.form.append('email', this.form.email)
      this.form.append('message', this.form.message)
    },

    async fetchAllContactForm () {
      await this.$store.dispatch('page/fetchAllContactForm')
    }
  }
}
</script>

<style lang="stylus" scoped>
#communication-contact-form textarea {
  height: 200px;
}
</style>
