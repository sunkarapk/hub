use crate::utils::error::Result;

use clap::Clap;

#[derive(Debug, Clap)]
#[clap(alias = "l")]
pub struct Labels {}

impl Labels {
    pub fn run(self) -> Result {
        Ok(())
    }
}
