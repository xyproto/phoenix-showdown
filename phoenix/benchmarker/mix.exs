defmodule Benchmarker.Mixfile do
  use Mix.Project

  def project do
    [ app: :benchmarker,
      version: "0.0.1",
      elixir: "~> 0.15.1",
      deps: deps ]
  end

  # Configuration for the OTP application
  def application do
    [
      mod: { Benchmarker, [] },
      applications: [:phoenix]
    ]
  end

  # Returns the list of dependencies in the format:
  # { :foobar, git: "https://github.com/elixir-lang/foobar.git", tag: "0.1" }
  #
  # To specify particular versions, regardless of the tag, do:
  # { :barbat, "~> 0.1", github: "elixir-lang/barbat" }
  defp deps do
    [
      {:phoenix, github: "chrismccord/phoenix"},
      {:cowboy, "~> 1.0.0"}
    ]
  end
end
