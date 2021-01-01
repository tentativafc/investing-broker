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
                    <h1 class="h5 mb-1">Bem vindo novamente!</h1>
                    <p class="text-muted mb-4">
                      Preencha seu email e senha para acesso ao aplicativo.
                    </p>

                    <ValidationError />
                    <!-- ValidationError Component -->

                    <form
                      class="user"
                      id="app"
                      @submit="validateAndSubmit"
                      novalidate="true"
                    >
                      <div class="form-group">
                        <input
                          type="email"
                          class="form-control form-control-user"
                          id="exampleInputEmail"
                          placeholder="Email"
                          v-model="email"
                        />
                      </div>
                      <div class="form-group">
                        <input
                          type="password"
                          class="form-control form-control-user"
                          id="exampleInputPassword"
                          placeholder="Senha"
                          v-model="password"
                        />
                      </div>
                      <a
                        href=""
                        class="btn btn-success btn-block"
                        v-on:click="validateAndSubmit"
                      >
                        Entrar
                      </a>
                    </form>

                    <div class="row mt-4">
                      <div class="col-12 text-center">
                        <p class="text-muted mb-2">
                          <a
                            href="/recover"
                            class="text-muted font-weight-medium ml-1"
                            >Esqueceu sua senha?</a
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
  <!-- end page -->
</template>

<script>
import { mapMutations } from 'vuex'
import { LOGIN } from '@/store/actions.type'
import { SET_VALIDATION_ERROR } from '@/store/mutations.type'
import { validEmail } from '@/common/functions'
import ValidationError from "./ValidationError.vue"

export default {
  name: "LoginComponent",
  components: {ValidationError},
  data() {
    return {
      email: null,
      password: null
    }
  },
  methods: {
    validateAndSubmit: function(e) {
      e.preventDefault()
      let errors = []
      if (!this.email) {
        errors.push('O e-mail é obrigatório.')
      } else if (!validEmail(this.email)) {
        errors.push('Utilize um e-mail válido.')
      }
      if (!this.password) {
        errors.push('A senha é obrigatória.')
      }
      if (errors.length) {
        this.[SET_VALIDATION_ERROR](errors)
      } else {
        this.$store
          .dispatch(LOGIN, { email: this.email, password: this.password })
          .then(() => this.$router.push({ name: 'home' }))
      }
      return false;
    },
    ...mapMutations([SET_VALIDATION_ERROR]),
  }
}
</script>
