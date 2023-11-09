package api

import (
	"github.com/gofiber/fiber/v2"
)

func LoginFormHandler(c *fiber.Ctx) error {
	if userCheck := c.Locals("user"); userCheck != nil {
		return c.SendString(`
		<form class="w-full max-w-sm flex flex-col gap-2 items-center text-center justify-center" hx-post="/api/token">
		<p id="hello">😊 Already Signed In! 😊</p>
		<p>Try going go this page to test it out!:</p>
		<a
			class="text-white bg-gradient-to-r from-purple-500 to-pink-500 hover:bg-gradient-to-l focus:ring-4 focus:outline-none focus:ring-purple-200 dark:focus:ring-purple-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center mr-2 mb-2"
			href="/restricted"
		>
			Shhh! Secret!
		</a>
		</form>`)
	} else {
		return c.SendString(`
	<form
		class="w-full max-w-sm flex flex-col gap-2 items-center text-center justify-center"
		hx-post="/api/token"
	>
		<div class="md:flex md:items-center">
		<div class="md:w-1/3">
			<label
			class="block font-bold md:text-right mb-1 md:mb-0 pr-4"
			for="user"
			>
			Username
			</label>
		</div>
		<div class="md:w-2/3">
			<input
			class="bg-gray-200 text-black appearance-none border-2 border-gray-200 rounded w-full leading-tight focus:outline-none focus:bg-white focus:border-purple-500"
			id="user"
			name="user"
			type="text"
			value="Danny DeVito"
			/>
		</div>
		</div>

		<div class="md:flex md:items-center">
		<div class="md:w-1/3">
			<label
			class="block font-bold md:text-right mb-1 md:mb-0 pr-4"
			for="pass"
			value="trashman69"
			>
			Password
			</label>
		</div>
		<div class="md:w-2/3">
			<input
			class="bg-gray-200 text-black appearance-none border-2 border-gray-200 rounded w-full leading-tight focus:outline-none focus:bg-white focus:border-purple-500"
			id="pass"
			name="pass"
			type="password"
			placeholder="******************"
			/>
		</div>
		</div>

		<button
		class="text-white bg-gradient-to-r from-purple-500 to-pink-500 hover:bg-gradient-to-l focus:ring-4 focus:outline-none focus:ring-purple-200 dark:focus:ring-purple-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center mr-2 mb-2"
		type="submit"
		>
		Sign In
		</button>
	</form>
	`)
	}
}
