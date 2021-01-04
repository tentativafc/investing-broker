<template>
  <div class="bg-primary">
    <div class="container">
      <div class="row">
        <div class="col-12">
          <div class="d-flex align-items-center min-vh-100">
            <div class="w-100 d-block bg-white shadow-lg rounded my-5">
              <div class="row">
                <div
                  class="col-lg-5 d-none d-lg-block bg-login rounded-left"
                ></div>
                <div class="col-lg-7">
                  <div class="p-5">
                    <div class="text-center">
                      <a href="index.html" class="d-block mb-5">
                        <img
                          src="assets/images/logo-dark.png"
                          alt="app-logo"
                          height="18"
                        />
                      </a>
                    </div>
                    <h1 class="h5 mb-1">Trocar Senha</h1>
                    <p class="text-muted mb-4">
                      Digite seu email e enviaremos um email com as instruções
                      para a renovação.
                    </p>

                    <div class="form-group">
                      <input
                        type="email"
                        class="form-control"
                        :class="{ 'is-invalid': $v.email.$error }"
                        id="email"
                        placeholder="Email"
                        v-model="email"
                      />
                      <div class="invalid-feedback" v-if="!$v.email.required">
                        Campo obrigatório
                      </div>
                      <div class="invalid-feedback" v-if="!$v.email.email">
                        Email inválido
                      </div>
                    </div>
                    <button
                      type="submit"
                      class="btn btn-success btn-block"
                      v-on:click="submit"
                    >
                      Recuperar senha
                    </button>

                    <div class="row mt-5">
                      <div class="col-12 text-center">
                        <p class="text-muted">
                          Já possui conta?
                          <a
                            href="/login"
                            class="text-muted font-weight-medium ml-1"
                            ><b>Login</b></a
                          >
                        </p>
                        <p class="text-muted mb-0">
                          Não possui conta?
                          <a
                            href="/register"
                            class="text-muted font-weight-medium ml-1"
                            ><b>Registre-se</b></a
                          >
                        </p>
                      </div>
                      <!-- end col -->
                    </div>
                    <!-- end row -->
                  </div>
                  <!-- end .padding-5 -->
                </div>
                <!-- end col -->
              </div>
              <!-- end row -->
            </div>
            <!-- end .w-100 -->
          </div>
          <!-- end .d-flex -->
        </div>
        <!-- end col-->
      </div>
      <!-- end row -->
    </div>
    <!-- end container -->
  </div>
</template>

<script>
import { required, email } from 'vuelidate/lib/validators'

import { FORGET_PASSWORD } from '@/store/actions.type'
export default {
  data() {
    return {
      email: null
    }
  },
  validations: {
    email: {
      required,
      email
    }
  },
  methods: {
    submit: function(e) {
      e.preventDefault()
      this.$v.$touch()
      if (!this.$v.$invalid) {
        this.$store
          .dispatch(FORGET_PASSWORD, { email: this.email })
          .then(() => this.$router.push({ name: 'login' }))
      }
    }
  }
}
</script>
