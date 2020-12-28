<template>
  <div class="bg-primary">
    <div class="container">
      <div class="row">
        <div class="col-12">
          <div class="d-flex align-items-center min-vh-100">
            <div class="w-100 d-block bg-white shadow-lg rounded my-5">
              <div class="row">
                <div
                  class="col-lg-5 d-none d-lg-block bg-register rounded-left"
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
                    <h1 class="h5 mb-1">Crie uma nova conta!</h1>
                    <p class="text-muted mb-4">
                      Não possui uma conta? Crie uma, leva menos de 1 minuto.
                    </p>

                    <Error />
                    <!-- Error Component -->

                    <form
                      class="user"
                      id="app"
                      @submit="validateAndSubmit"
                      novalidate="true"
                    >
                      <div class="form-group row">
                        <div class="col-sm-6 mb-3 mb-sm-0">
                          <input
                            type="text"
                            class="form-control form-control-user"
                            id="firstName"
                            v-model="firstname"
                            placeholder="Primeiro nome"
                          />
                        </div>
                        <div class="col-sm-6">
                          <input
                            type="text"
                            class="form-control form-control-user"
                            id="lastName"
                            v-model="lastname"
                            placeholder="Último nome"
                          />
                        </div>
                      </div>
                      <div class="form-group">
                        <input
                          type="email"
                          class="form-control form-control-user"
                          id="email"
                          v-model="email"
                          placeholder="Email"
                        />
                      </div>
                      <div class="form-group row">
                        <div class="col-sm-6 mb-3 mb-sm-0">
                          <input
                            type="password"
                            class="form-control form-control-user"
                            v-model="password"
                            id="password"
                            placeholder="Senha"
                          />
                        </div>
                        <div class="col-sm-6">
                          <input
                            type="password"
                            class="form-control form-control-user"
                            v-model="repeat_password"
                            id="repeatPassword"
                            placeholder="Confirmação de senha"
                          />
                        </div>
                      </div>
                      <a
                        href=""
                        class="btn btn-success btn-block"
                        v-on:click="validateAndSubmit"
                      >
                        Criar conta
                      </a>
                    </form>

                    <div class="row mt-4">
                      <div class="col-12 text-center">
                        <p class="text-muted mb-0">
                          Já possui uma conta?
                          <a
                            href="/login"
                            class="text-muted font-weight-medium ml-1"
                            ><b>Entre</b></a
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
import { REGISTER } from '@/store/actions.type'
import { SET_ERROR } from '@/store/mutations.type'
import { validEmail } from '@/common/functions'
import Error from './Error.vue'

export default {
  name: 'RegisterComponent',
  components: { Error },
  data() {
    return {
      firstname: null,
      lastname: null,
      email: null,
      password: null,
      repeat_password: null
    }
  },
  methods: {
    validateAndSubmit: function(e) {
      e.preventDefault()

      let errors = []
      if (!this.firstname) {
        errors.push('O primeiro nome é obrigatório.')
      }
      if (!this.lastname) {
        errors.push('O último nome é obrigatório.')
      }
      if (!this.email) {
        errors.push('O e-mail é obrigatório.')
      } else if (!validEmail(this.email)) {
        errors.push('Utilize um e-mail válido.')
      }
      if (!this.password) {
        errors.push('A senha é obrigatória.')
      }

        if (!this.repeat_password) {
                errors.push('A conformação de senha é obrigatória.')
            }


        if (this.password !== this.repeat_password) {
                errors.push('A senha e conformação são diferentes.')
        }

      if (errors.length) {
        this.[SET_ERROR](errors)
      } else {
        this.$store
          .dispatch(REGISTER, { firstname: this.firstname, lastname:this.lastname, email: this.email, password: this.password })
          .then(() => this.$router.push({ name: 'home' }))
      }
      return false;
    },
    ...mapMutations([SET_ERROR]),
  }
}
</script>
